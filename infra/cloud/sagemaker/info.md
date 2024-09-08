# Information


## Relavant Links

[Deploy and manage machine learning pipelines with Terraform using Amazon SageMaker](https://aws.amazon.com/blogs/machine-learning/deploy-and-manage-machine-learning-pipelines-with-terraform-using-amazon-sagemaker/)

[GitHub Repo for above link](https://github.com/aws-samples/amazon-sagemaker-ml-pipeline-deploy-with-terraform)

[ML-related code for training and inference with a Docker image](https://github.com/aws/amazon-sagemaker-examples/tree/main/advanced_functionality/scikit_bring_your_own)

[Amazon SageMaker Domain in VPC only mode to support SageMaker Studio with auto shutdown Lifecycle Configuration and SageMaker Canvas with Terraform](https://aws.amazon.com/blogs/machine-learning/amazon-sagemaker-domain-in-vpc-only-mode-to-support-sagemaker-studio-with-auto-shutdown-lifecycle-configuration-and-sagemaker-canvas-with-terraform/)


### This directory contains the pipeline for connecting sagemaker to the infrastructure being developed.

Necessary items:
 - An existing VPC
 - Existing subnets within the VPC
 - An S3 bucket for storing data
 - IAM roles and policies for SageMaker access