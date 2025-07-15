resource "aws_s3_bucket" "bucket_imagem_perfil" {
  bucket = "mcastegnaro-bucket-imagem-perfil"
}

resource "aws_s3_bucket" "bucket_imagem_pets" {
  bucket = var.bucket_imagem_pets
}

