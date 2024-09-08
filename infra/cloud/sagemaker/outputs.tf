# Outputs
output "data_eng_ml_notebook" {
  description = "The name of the SageMaker notebook instance"
  value       = aws_sagemaker_notebook_instance.data-eng-ml-notebook.name
}

output "sagemaker_role" {
  description = "sagemaker_role"
  value       = aws_iam_role.sagemaker_role.arn
}