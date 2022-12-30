terraform {
  required_providers {
    alma = {
      source = "cnm.team/cenomi/alma"
    }
  }
}


provider "alma" {
  host = "http://service-catalog.platform.k8s.dev.cnm.team"
}

data "alma_services" "all" {}

data "alma_service" "default" {
  name = "order"
}

output "service_metadata" {
  value = data.alma_service.default
}
