variable "env" {}
variable "alb_terget_group_arn" {}
variable "api_alb_terget_group_arn" {}
variable "pri_sub_0_id" {}
variable "pri_sub_1_id" {}
variable "security_group_80_id" {}
variable "security_group_8080_id" {}
variable "container_definitions" {}
variable "self" {}

resource "aws_ecs_cluster" "cluster" {
  name = "FTRA-cluster"
  tags = {
    Name = "FTRA-cluster"
  }
}

resource "aws_ecs_task_definition" "ecs_task_definition" {
  family                   = "FTRA-datadog-task"
  memory                   = "1024"
  network_mode             = "awsvpc"
  cpu                      = "512"
  requires_compatibilities = ["FARGATE"]
  container_definitions    = var.container_definitions
  task_role_arn            = "arn:aws:iam::${var.self}:role/ecsTaskExecutionRole"
  execution_role_arn       = "arn:aws:iam::${var.self}:role/ecsTaskExecutionRole"
}

resource "aws_ecs_service" "ecs_service" {
  name            = "FTRA-datadog-service"
  cluster         = aws_ecs_cluster.cluster.id
  task_definition = aws_ecs_task_definition.ecs_task_definition.arn
  desired_count   = 2
  launch_type     = "FARGATE"
  load_balancer {
    target_group_arn = var.alb_terget_group_arn
    container_name   = "FTRA-UI-container"
    container_port   = 80
  }

  load_balancer {
    target_group_arn = var.api_alb_terget_group_arn
    container_name   = "FTRA-container"
    container_port   = 8080
  }

  network_configuration {
    assign_public_ip = true

    subnets = [
      var.pri_sub_0_id,
      var.pri_sub_1_id
    ]

    security_groups = [
      var.security_group_80_id,
      var.security_group_8080_id
    ]
  }
}
