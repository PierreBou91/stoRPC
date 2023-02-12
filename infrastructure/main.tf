terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.54.0"
    }
    docker = {
      source  = "kreuzwerker/docker"
      version = "3.0.1"
    }
  }
}


# Configure the AWS Provider
provider "aws" {
  region = "eu-west-3"
}

# Reference the Docker image
data "docker_registry_image" "stoRPC_server" {
  name = "pierrebou91/storpc-server:latest"
}

resource "docker_image" "stoRPC_server" {
  name          = data.docker_registry_image.stoRPC_server.name
  pull_triggers = [data.docker_registry_image.stoRPC_server.sha256_digest]
}

# AWS ECS Cluster
resource "aws_ecs_cluster" "stoRPC" {
  name = "stoRPC"
}

resource "aws_ecs_task_definition" "server_task" {
  family = "server_task_family"
  container_definitions = jsonencode([
    {
      name  = "stoRPC_server"
      image = "${resource.docker_image.stoRPC_server.name}"
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
    }
  ])
}

resource "aws_ecs_service" "server_service" {
  name            = "server_service"
  task_definition = aws_ecs_task_definition.server_task.arn
  cluster         = aws_ecs_cluster.stoRPC.id
  desired_count   = 1
}

