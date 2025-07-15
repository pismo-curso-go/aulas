output "bucket_id_imagem_perfil" {
  description = "ID de acesso ao bucket criado na AWS"
  value       = aws_s3_bucket.bucket_imagem_perfil.id
}

output "bucket_arn_imagem_perfil" {
  description = "ARN de acesso ao bucket criado na AWS"
  value       = aws_s3_bucket.bucket_imagem_perfil.arn
}
