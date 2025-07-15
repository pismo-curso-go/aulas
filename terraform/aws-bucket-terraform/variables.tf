variable "location" {
  description = "Região onde os recursos serão criados na aws"
  type        = string
  default     = "sa-east-1"
}

variable "bucket_imagem_pets" {
  description = "Nome do bucket de upload de imagens de pets"
  type        = string
  default     = "mcastegnaro-bucket-imagem-pets"
}
