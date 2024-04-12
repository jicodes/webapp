packer {
  required_plugins {
    googlecompute = {
      version = ">= 1.0.0, < 2.0.0"
      source  = "github.com/hashicorp/googlecompute"
    }
  }
}

variable "project_id" {
  type    = string
  default = "true-server-412502"
}

variable "source_image_family" {
  type    = string
  default = "centos-stream-8"
}

variable "zone" {
  type    = string
  default = "us-central1-a"
}

variable "disk_size" {
  type    = number
  default = 20
}

variable "disk_type" {
  type    = string
  default = "pd-standard"
}

variable "app_path" {
  type    = string
  default = "webapp"
}

variable "app_properties_path" {
  type    = string
  default = "app.properties"
}

variable "service_file_path" {
  type    = string
  default = "webapp.service"
}

variable "opsagent_config_file_path" {
  type    = string
  default = "config.yaml"
}

source "googlecompute" "csye6225-app-custom-image" {
  project_id              = var.project_id
  source_image_family     = var.source_image_family
  zone                    = var.zone
  disk_size               = var.disk_size
  disk_type               = var.disk_type
  image_name              = "csye6225-{{timestamp}}"
  image_family            = "csye6225-app-image"
  image_storage_locations = ["us"]
  ssh_username            = "packer"
}

build {
  name = "packer"
  sources = [
    "source.googlecompute.csye6225-app-custom-image"
  ]

  provisioner "shell" {
    script = "updateOS.sh"
  }

  provisioner "shell" {
    script = "setupGolang.sh"
  }

  // provisioner "shell" {
  //   script = "setupDB.sh"
  // }

  provisioner "shell" {
    script = "setupVmUser.sh"
  }

  provisioner "shell" {
    inline = [
      "curl -sSO https://dl.google.com/cloudagents/add-google-cloud-ops-agent-repo.sh",
      "sudo bash add-google-cloud-ops-agent-repo.sh --also-install"
    ]
  }

  provisioner "file" {
    source      = var.opsagent_config_file_path
    destination = "/tmp/config.yaml"
  }

  provisioner "shell" {
    inline = [
      "sudo mv /tmp/config.yaml /etc/google-cloud-ops-agent/config.yaml",
      "sudo systemctl restart google-cloud-ops-agent"
    ]
  }

  provisioner "shell" {
    script = "setupAppDir.sh"
  }

  provisioner "file" {
    source      = var.app_path
    destination = "/opt/myapp/webapp"
  }

  provisioner "file" {
    source      = var.app_properties_path
    destination = "/opt/myapp/app.properties"
  }

  provisioner "shell" {
    script = "setupAppPermission.sh"
  }

  provisioner "shell" {
    inline = [
      # Change systemd directory permissions
      "sudo chown packer:packer /etc/systemd/system/",
      "sudo chmod 755 /etc/systemd/system/"
    ]
  }

  provisioner "file" {
    source      = var.service_file_path
    destination = "/etc/systemd/system/webapp.service"
  }

  provisioner "shell" {
    script = "setupAppService.sh"
  }
}