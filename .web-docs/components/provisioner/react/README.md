  Include a short description about the provisioner. This is a good place
  to call out what the provisioner does, and any additional text that might
  be helpful to a user. See https://www.packer.io/docs/provisioner/null
-->

The `react` provisioner is used to install a compiled React-based APP in AWS AMI image


<!-- Provisioner Configuration Fields -->

**Required**

- `distSource` (string) - The path to a local dist file to upload to the machine. The path can be absolute or relative.
   If it is relative, it is relative to the working directory when Packer is executed.
- `appDomain` (string) - the SSL-enabled domain that will serve the deployed HTTP React APP instance.
- `sslCertBase64` (string) - is a __base64 encoded__ string of the content of SSL certificate file for the SSL-enabled
  domain, for example `app.mycompany.com` given the `appDomain` is `app.mycompany.com`.
- `sslCertKeyBase64` (string) - is a __base64 encoded__ string of the content of SSL certificate key file for the SSL
  enabled domain, for example `app.mycompany.com` given the `appDomain` is `app.mycompany.com`.


<!--
  Optional Configuration Fields

  Configuration options that are not required or have reasonable defaults
  should be listed under the optionals section. Defaults values should be
  noted in the description of the field
-->

**Optional**

- `nodeVersion` (string) - The Node.js version running the React app; default to "18"
- `homeDir` (string) - The `$Home` directory in AMI image; default to `/home/ubuntu`

<!--
  A basic example on the usage of the provisioner. Multiple examples
  can be provided to highlight various configurations.

-->

### Example Usage

```hcl
packer {
  required_plugins {
    amazon = {
      version = ">= 0.0.2"
      source  = "github.com/hashicorp/amazon"
    }
    qubitpi = {
      version = ">= 0.0.50"
      source = "github.com/QubitPi/qubitpi"
    }
  }
}

source "amazon-ebs" "qubitpi" {
  ami_name              = "packer-plugin-qubitpi-acc-test-ami-react"
  force_deregister      = "true"
  force_delete_snapshot = "true"

  instance_type = "t2.micro"
  launch_block_device_mappings {
    device_name           = "/dev/sda1"
    volume_size           = 8
    volume_type           = "gp2"
    delete_on_termination = true
  }
  region = "us-west-1"
  source_ami_filter {
    filters = {
      name                = "ubuntu/images/*ubuntu-*-22.04-amd64-server-*"
      root-device-type    = "ebs"
      virtualization-type = "hvm"
    }
    most_recent = true
    owners      = ["099720109477"]
  }
  ssh_username = "ubuntu"
}

build {
  sources = [
    "source.amazon-ebs.qubitpi"
  ]

  provisioner "qubitpi-react-provisioner" {
    distSource       = "./dist"
    homeDir          = "/home/ubuntu"
    sslCertBase64    = "YXNkZnNnaHRkeWhyZXJ3ZGZydGV3ZHNmZ3RoeTY0cmV3ZGZyZWd0cmV3d2ZyZw=="
    sslCertKeyBase64 = "MzI0NXRnZjk4dmJoIGNsO2VbNDM1MHRdzszNDM1b2l0cmo="
    appDomain        = "app.mycompany.com"
  }
}
```
