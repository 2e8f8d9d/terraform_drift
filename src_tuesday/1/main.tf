terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.46.0"
    }
  }
}

# Configure the Microsoft Azure Provider
provider "azurerm" {
  features {}

  subscription_id = "99c19e8a-0929-4a39-af63-a232d1686654"
}

resource "azurerm_resource_group" "my_resource_group" {
    name = "terraform-rg"
    location = "centralus"
}