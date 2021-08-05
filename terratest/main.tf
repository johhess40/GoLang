/*
This code is being used to test automated unit/integration testing with Terratest
*/

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "hubspokerg" {
  name     = var.iac_test_rg.name
  location = var.iac_test_rg.location
  tags     = var.iac_test_rg.tags
}

resource "azurerm_virtual_network" "vnets" {
  for_each            = var.hub_and_spokes
  address_space       = each.value["address_space"]
  location            = each.value["location"]
  name                = each.value["name"]
  resource_group_name = azurerm_resource_group.hubspokerg.name

  dynamic "subnet" {
    for_each = each.value["subnets"]
    content {
      name           = subnet.value["name"]
      address_prefix = subnet.value["address_prefix"]
    }
  }
}

resource "azurerm_public_ip" "fwpip" {
  for_each            = var.fw_pips
  allocation_method   = each.value["allocation_method"]
  location            = each.value["location"]
  name                = each.value["name"]
  resource_group_name = azurerm_resource_group.hubspokerg.name
}

resource "azurerm_firewall" "hsfw" {
  for_each            = var.hs_fws
  name                = each.value["name"]
  resource_group_name = azurerm_resource_group.hubspokerg.name
  location            = each.value["location"]
  dynamic "ip_configuration" {
    for_each = each.value["ip_configuration"]
    content {
      name                 = ip_configuration.value["name"]
      public_ip_address_id = ip_configuration.value["public_ip_address_id"]
    }
  }
}