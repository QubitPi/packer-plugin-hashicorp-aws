# Copyright (c) Jiaqi
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
  }
}

source "docker" "hashistack" {
  image  = "jack20191124/packer-plugin-hashistack-acc-test-base:latest"
  discard = true
}

build {
  sources = [
    "source.docker.hashistack"
  ]

  provisioner "hashistack-kong-api-gateway-provisioner" {
    homeDir              = "/"
    sslCertBase64        = "YXNkZnNnaHRkeWhyZXJ3ZGZydGV3ZHNmZ3RoeTY0cmV3ZGZyZWd0cmV3d2ZyZw=="
    sslCertKeyBase64     = "MzI0NXRnZjk4dmJoIGNsO2VbNDM1MHRdzszNDM1b2l0cmo="
    kongApiGatewayDomain = "api.mycompany.com"
  }
}
