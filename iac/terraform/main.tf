
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
