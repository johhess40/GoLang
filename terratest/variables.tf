variable "iac_test_rg" {
  type = object({
    name     = string
    location = string
    tags = object({
      Env = string
    })
  })
}

variable "hub_and_spokes" {
  type = map(object({
    address_space = list(string)
    name          = string
    location      = string
    subnets = map(object({
      name           = string
      address_prefix = string
    }))
  }))
}

variable "fw_pips" {
  type = map(object({
    name = string
    location = string
    allocation_method = string
  }))
}
variable "hs_fws" {
  type = map(object({
    name = string
    location = string
    ip_configuration = map(object({
      name = string
      public_ip_address_id = string
    }))
  }))
}