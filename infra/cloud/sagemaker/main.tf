provider "aws" {
  region  = var.region
}

# Create a SageMaker Model
resource "aws_sagemaker_model" "data-eng-ml-model" {
  name          = "data-eng-ml-model"
  execution_role_arn = aws_iam_role.sagemaker_role.arn

  primary_container {
    image = data.aws_sagemaker_prebuilt_ecr_image.test.registry_path 
  }

  vpc_config {
    subnets = var.subnet_ids
    security_group_ids = [] # optional: Add security group IDs
  }

  tags = {
    Name = "data-eng-ml-model"
  }
}

# Create an IAM role for SageMaker
resource "aws_iam_role" "sagemaker_role" {
  name               = "sagemaker_role"
  assume_role_policy = data.aws_iam_policy_document.sagemaker_assume_role.json

}

data "aws_iam_policy_document" "sagemaker_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}

#  Attach policies to the IAM role
resource "aws_iam_role_policy_attachment" "sagemaker_full_access" {
  role       = aws_iam_role.sagemaker_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSageMakerFullAccess"
}

resource "aws_iam_role_policy_attachment" "s3_access" {
  role       = aws_iam_role.sagemaker_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

# Create a SageMaker Notebook Instance
resource "aws_sagemaker_notebook_instance" "data-eng-ml-notebook" {
  name                  = "data-eng-ml-notebook"
  instance_type         = "ml.t2.medium"
  role_arn              = aws_iam_role.sagemaker_role.arn
  subnet_id             = var.subnet_ids[0] # Use the first subnet in the list
  security_groups       = [""] # Add security group IDs if necessary

  tags = {
    Name = "data-eng-ml-notebook"
  }
}

data "aws_sagemaker_prebuilt_ecr_image" "test" {
  repository_name = "image-classification"
}