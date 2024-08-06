
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = ">= 3.7.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "fl-rg"
    storage_account_name = "flterraformsa"
    container_name       = "terraform"
    key                  = "terraform.tfstate"
    use_oidc = true
  }
}

provider "azurerm" {
  features {}
  use_oidc = true
}

resource "azurerm_resource_group" "test-rg" {
  name     = "fullstack-renamed"
  location = "West US"
}

resource "azurerm_container_registry" "acr" {
  name                = "fullstackloadgen"
  resource_group_name  = azurerm_resource_group.test-rg.name
  location            = azurerm_resource_group.test-rg.location
  sku                 = "Basic"

  admin_enabled       = true
}

resource "azurerm_kubernetes_cluster" "k8s" {
  location            = azurerm_resource_group.test-rg.location
  name                = "fullstackloadgen"
  resource_group_name = azurerm_resource_group.test-rg.name
  dns_prefix          = "fullstackloadgen"

  identity {
    type = "SystemAssigned"
  }

  default_node_pool {
    name       = "agentpool"
    vm_size    = "Standard_B2s"
    node_count = 2
  }

  network_profile {
    network_plugin    = "kubenet"
    load_balancer_sku = "standard"
  }
}

resource "azurerm_role_assignment" "k8srole" {
  principal_id                     = azurerm_kubernetes_cluster.k8s.kubelet_identity[0].object_id
  role_definition_name             = "AcrPull"
  scope                            = azurerm_container_registry.acr.id
  skip_service_principal_aad_check = true
}

resource "azurerm_cosmosdb_account" "cosmos" {
  name                      = "fullstackloadgen"
  location                  = azurerm_resource_group.test-rg.location
  resource_group_name       = azurerm_resource_group.test-rg.name
  offer_type                = "Standard"
  kind                      = "GlobalDocumentDB"
  enable_automatic_failover = false

  geo_location {
    location          = azurerm_resource_group.test-rg.location
    failover_priority = 0
  }

  consistency_policy {
    consistency_level       = "BoundedStaleness"
    max_interval_in_seconds = 300
    max_staleness_prefix    = 100000
  }

  depends_on = [azurerm_resource_group.test-rg]
}

data "azurerm_client_config" "current" {}

resource "azurerm_key_vault" "akv" {
  name                = "fullstackloadgen"
  location            = azurerm_resource_group.test-rg.location
  resource_group_name = azurerm_resource_group.test-rg.name
  sku_name            = "standard"
  tenant_id           = data.azurerm_client_config.current.tenant_id
  enable_rbac_authorization = true
}

resource "azurerm_role_assignment" "akv_sp" {
  scope                = azurerm_key_vault.akv.id
  role_definition_name = "Key Vault Secrets Administrator"
  principal_id         = data.azurerm_client_config.current.object_id
}

resource "azurerm_key_vault_secret" "cosmosdb_connection_string" {
  name         = "CosmosDBConnectionString"
  value        = azurerm_cosmosdb_account.cosmos.connection_strings[0]
  key_vault_id = azurerm_key_vault.akv.id
}