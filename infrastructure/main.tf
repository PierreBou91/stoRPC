terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.54.0"
    }
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "eu-west-3"
}

# Create VPC
resource "aws_vpc" "stoRPC_vpc" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name        = "stoRPC_vpc"
    Terraform   = "true"
    Environment = "dev"
    StoRPC      = "true"
  }
}

# Create Subnet
resource "aws_subnet" "stoRPC_server_subnet" {
  depends_on = [
    aws_vpc.stoRPC_vpc
  ]
  vpc_id     = aws_vpc.stoRPC_vpc.id
  cidr_block = "10.0.1.0/24"

  tags = {
    Name        = "server_subnet"
    Terraform   = "true"
    Environment = "dev"
    StoRPC      = "true"
  }
}

resource "aws_security_group" "stoRPC_security_group" {
  name        = "allow_internet_and_TCP_8080"
  description = "Allow internet and TCP 8080"
  vpc_id      = aws_vpc.stoRPC_vpc.id

  ingress {
    description = "TLS from VPC"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = [aws_subnet.stoRPC_server_subnet.cidr_block]
  }

  #   ingress {
  #     description = "Basic HTTP from VPC"
  #     from_port   = 80
  #     to_port     = 80
  #     protocol    = "tcp"
  #     cidr_blocks = [aws_subnet.stoRPC_server_subnet.cidr_block]
  #   }

  #   ingress {
  #     description = "Service port from VPC"
  #     from_port   = 8080
  #     to_port     = 8080
  #     protocol    = "tcp"
  #     cidr_blocks = [aws_subnet.stoRPC_server_subnet.cidr_block]
  #   }

  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_network_acl" "stoRPC_network_acl" {
  vpc_id = aws_vpc.stoRPC_vpc.id

  egress {
    protocol   = "-1"
    rule_no    = 200
    action     = "allow"
    cidr_block = aws_subnet.stoRPC_server_subnet.cidr_block
    from_port  = 0
    to_port    = 0
  }

  ingress {
    protocol   = "-1"
    rule_no    = 100
    action     = "allow"
    cidr_block = aws_subnet.stoRPC_server_subnet.cidr_block
    from_port  = 0
    to_port    = 0
  }

  tags = {
    Name = "main"
  }
}

# Reference the Docker image
# data "docker_registry_image" "stoRPC_server_registry" {
#   name = "pierrebou91/storpc-server:latest" # Make this a variable
# }

# resource "docker_image" "stoRPC_server_image" {
#   #   depends_on = [
#   #     data.docker_registry_image.stoRPC_server_registry
#   #   ]
#   #   name          = data.docker_registry_image.stoRPC_server_registry.name
#   name = "pierrebou91/storpc-server:latest"
#   #   pull_triggers = [data.docker_registry_image.stoRPC_server_registry.sha256_digest]
# }

# AWS ECS Cluster
resource "aws_ecs_cluster" "stoRPC_ecs" {
  name = "stoRPC_ecs"
  tags = {
    Name        = "stoRPC_ecs"
    Terraform   = "true"
    Environment = "dev"
    StoRPC      = "true"
  }
}

resource "aws_ecs_task_definition" "server_task" {
  network_mode = "awsvpc"
  depends_on = [
    aws_ecs_cluster.stoRPC_ecs
  ]
  family                   = "server_task_family"
  memory                   = 512
  cpu                      = 256
  requires_compatibilities = ["FARGATE"]
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name  = "stoRPC_server"
      image = "pierrebou91/storpc-server:latest"
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
    }
  ])
  tags = {
    Name        = "server_task"
    Terraform   = "true"
    Environment = "dev"
    StoRPC      = "true"
  }
}

resource "aws_ecs_service" "server_service" {
  depends_on = [
    aws_ecs_cluster.stoRPC_ecs,
    aws_ecs_task_definition.server_task
  ]
  name            = "server_service"
  task_definition = aws_ecs_task_definition.server_task.arn
  cluster         = aws_ecs_cluster.stoRPC_ecs.id
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets          = [aws_subnet.stoRPC_server_subnet.id]
    assign_public_ip = true
    security_groups  = [aws_security_group.stoRPC_security_group.id]
  }
  tags = {
    Name        = "server_service"
    Terraform   = "true"
    Environment = "dev"
    StoRPC      = "true"
  }
}


resource "aws_iam_role" "ecs_task_execution_role" {
  name = "ecs_task_execution_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action = "sts:AssumeRole",
        Effect = "Allow",
        Principal = {
          Service = "ecs-tasks.amazonaws.com"
        }
      }
    ]
  })
  tags = {
    Name        = "ecs_task_execution_role"
    Terraform   = "true"
    Environment = "dev"
    StoRPC      = "true"
  }
}
