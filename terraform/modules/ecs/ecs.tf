variable "env" {}
variable "aws_lb_target_group_arn" {}
variable "iam_role_arn" {}

resource "aws_ecs_cluster" "cluster" {
  name = "FTRA-cluster"
  tags = {
    Name = "FTRA-cluster"
  }
}

resource "aws_ecs_task_definition" "ecs_task_definition" {
  family                   = "service"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions    = "${file("task-definitions/service.json")}"

  placement_constraints {
    type       = "memberOf"
    expression = "attribute:ecs.availability-zone in [ap-northeast-1a, ap-northeast-1c]"
  }
}

resource "aws_ecs_service" "app_with_mackerel" {
  name            = "FTRA-service"
  cluster         = "${aws_ecs_cluster.cluster.id}"
  task_definition = "${aws_ecs_task_definition.ecs_task_definition.arn}"
  desired_count   = 3
  iam_role        = "${var.iam_role_arn}"
  depends_on      = ["aws_iam_role_policy.foo"]

  load_balancer {
    target_group_arn = "${var.aws_lb_target_group_arn}"
    container_name   = "app_with_mackerel"
    container_port   = 80
  }

  placement_constraints {
    type       = "memberOf"
    expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
  }
}
