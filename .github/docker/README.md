Plugin Acceptance Tests Base Image
==================================

[![badge](https://img.shields.io/badge/DockerHub-2596EC?style=for-the-badge&logo=docker&logoColor=white)](https://hub.docker.com/r/jack20191124/packer-plugin-qubitpi-acc-test-base)

This image is used by [packer-plugin-qubitpi](https://github.com/QubitPi/packer-plugin-qubitpi) to perform
acceptance tests in Docker without the need to test against AWS, which
[requires a real AWS credentials](https://developer.hashicorp.com/packer/tutorials/aws-get-started/aws-get-started-build-image#authenticate-to-aws)

To use this image, the test fixture can be 

```
packer {
  required_plugins {
    docker = {
      version = ">= 0.0.7"
      source  = "github.com/hashicorp/docker"
    }
  }
}

source "docker" "qubitpi" {
  image  = "jack20191124/packer-plugin-qubitpi-acc-test-base:latest"
  discard = true
}

build {
  sources = [
    "source.docker.qubitpi"
  ]

  provisioner "qubitpi-webservice-provisioner" {
    homeDir   = "/"
    warSource = "my-webservice.war"
  }
}
```

In this example, we are testing a [webservice provisioner](../../docs/provisioners/webservice.mdx) in Docker
