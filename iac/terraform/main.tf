
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