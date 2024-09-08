# Cloud Infrastructure Internship Project - ChakraTech (Summer 2024)

This repository showcases my contributions to ChakraTech's cloud infrastructure during my 8-week internship in the summer of 2024. It represents a snapshot of the work I completed before departing in August 2024.
Project Overview

During my internship at ChakraTech, a startup company, I worked on developing cloud infrastructure components using Terraform and Golang. This project demonstrates my skills in cloud computing, infrastructure as code, AWS services, and backend development.

## Important Note
This repository reflects my personal contributions up to August 2024. As other interns continued working on the project after my offboarding process, this codebase may not represent the final state of the project and may not be fully functional as a standalone system.
Project Structure and Technologies

## The project is organized into several main components:
### Data Platform

data-platform/

api/: API implementation using Golang

cmd/: Command-line tools for various data platform operations

config/: Configuration files for the data platform

css/: Stylesheets for any web interfaces

internal/: Internal packages used across the data platform

pkg/: Reusable Golang packages

public/: Public assets for web interfaces

routes/: API route definitions

videoUploads/: Handling video upload functionality

***Key features:***

Golang-based API for data processing and management

Integration with AWS services using AWS SDK for Go

Command-line tools for data platform management

### Infrastructure (infra)

infra/

.venv/: Python virtual environment for any Python-based tools

cloud/: Cloud infrastructure code

acm/: AWS Certificate Manager configurations

application-load-balancer/: Load balancer setup using Terraform

ec2/: EC2 instance configurations

enc/: Encryption-related code and configurations

rds/: RDS database setup using Terraform

s3_server/: S3 server configurations

s3_storage/: S3 storage setup using Terraform

sagemaker/: SageMaker configurations for machine learning tasks

scripts/: Infrastructure management scripts

vpn/: VPN setup for secure access

***Key features:***

Terraform configurations for provisioning and managing AWS resources

Golang connection files to interact with AWS services

Infrastructure as Code (IaC) approach for maintainable and versioned infrastructure

### Onboarding Setup

onboarding-setup/

git.md: Git workflow and best practices

aws-setup.md: Guide for setting up AWS CLI and permissions

terraform-intro.md: Introduction to Terraform and its usage in the project

golang-guide.md: Golang setup and coding standards

IT-CONTRIB.md: Contribution guidelines and code review process

README.md: Overview of the onboarding process

***Key features:***

To provide comprehensive guidance for new interns on tools, technologies, and workflows used in the project.

### Portal

portal/

platform/: Core platform code

node_modules/: Node.js dependencies

public/: Public assets for the web portal

src/: Source code for the web application

project/: Project-specific customizations and features

***Key features:***

Web-based interface for interacting with the data platform

Integration with backend APIs

User authentication and authorization

## My Contributions
During my 8-week internship, I focused on:

Setting up the initial project structure and Terraform configurations

Implementing core infrastructure components using AWS services and Terraform

Developing parts of the data platform API using Golang

Creating Golang connection files to interact with AWS services

Contributing to the portal's platform and project-specific code

Developing onboarding documentation to facilitate knowledge transfer

***Running the Project***
While this repository contains Makefiles and Dockerfiles, please note that running the entire system may not be possible due to:

Subsequent changes made by other interns after my departure

Potential missing components or dependencies not included in this repository

Possible changes in the overall architecture or design of the system

The need for specific AWS credentials and access keys

The primary purpose of this repository is to demonstrate the code and infrastructure I worked on during my internship.

## Skills Demonstrated
This project showcases my experience with:

Cloud infrastructure design and implementation

AWS services (EC2, S3, RDS, SageMaker, etc.)

Infrastructure as Code using Terraform

Golang development for backend services and AWS interactions

API development

Docker containerization

Makefile usage for project management

Technical documentation and onboarding material creation

## Contact
For any questions regarding my contributions or to discuss this project further, please feel free to contact me at thanhdatvu.203@gmail.com.
