
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
  features {
      key_vault {
        purge_soft_deleted_secrets_on_destroy = true
    }
  }
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

  key_vault_secrets_provider {
    secret_rotation_enabled = true
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

resource "random_pet" "db_username" {
  length = 2
}

resource "random_password" "db_password" {
  length  = 16
  special = true
  upper   = true
  lower   = true
}

resource "azurerm_postgresql_flexible_server" "main" {
  name                   = "fullstackloadgen"
  resource_group_name    = azurerm_resource_group.test-rg.name
  location               = azurerm_resource_group.test-rg.location
  administrator_login    = random_pet.db_username.id
  administrator_password = random_password.db_password.result

  authentication {
    active_directory_auth_enabled = "true"
    password_auth_enabled         = "true"
    tenant_id                     = data.azurerm_client_config.current.tenant_id
  }

  backup_retention_days        = 7
  geo_redundant_backup_enabled = "true"

  sku_name   = "B_Standard_B1ms"
  storage_mb = 65536
  version    = 11
}

resource "azurerm_postgresql_flexible_server_firewall_rule" "main" {
  name                = "fullstackloadgen"
  server_id           = azurerm_postgresql_flexible_server.main.id
  start_ip_address    = "0.0.0.0"
  end_ip_address      = "0.0.0.0"
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
  role_definition_name = "Key Vault Administrator"
  principal_id         = data.azurerm_client_config.current.object_id
}

resource "azurerm_role_assignment" "akv_sp_k8s" {
  principal_id                     = azurerm_kubernetes_cluster.k8s.kubelet_identity[0].object_id
  role_definition_name             = "Key Vault Secrets User"
  scope                            = azurerm_key_vault.akv.id
}

resource "azurerm_key_vault_secret" "db_login" {
  name         = "db-login"
  value        = random_pet.db_username.id
  key_vault_id = azurerm_key_vault.akv.id
}

resource "azurerm_key_vault_secret" "db_password" {
  name         = "db-password"
  value        = random_password.db_password.result
  key_vault_id = azurerm_key_vault.akv.id
}