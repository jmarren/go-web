
variable "vpc_cidr_block" {
  type        = string
  description = "CIDR block for  vpc"
  default     = "10.0.0.0/16"
}


variable "app_name" {
  type        = string
  default     = "myapp"
  description = "The name of the application"
}

