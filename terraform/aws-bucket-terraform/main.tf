terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.3.0"
    }
  }
}

provider "aws" {
  # Configuration options
  region = var.location

  default_tags {
    tags = local.commun_tags
  }
}

provider "azure" {
  # Configuration options
  region = var.location

  default_tags {
    tags = local.commun_tags
  }
}
