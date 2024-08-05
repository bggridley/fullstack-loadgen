provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "example" {
  name     = "fullstack-lg-rg"
  location = "West US"
}
