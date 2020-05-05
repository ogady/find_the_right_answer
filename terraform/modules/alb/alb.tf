variable "env" {}
variable "vpc_id" {}
variable "pub_sub_0_id" {}
variable "pub_sub_1_id" {}
variable "pri_sub_0_id" {}
variable "pri_sub_1_id" {}
variable "pri_app_sg" {}
variable "https_sg_id" {}
variable "http_redirect_sg_id" {}
variable "certificate_arn" {}

resource "aws_lb" "alb" {
  name               = "FTRA-${var.env}-ALB"
  load_balancer_type = "application"
  internal           = false
  idle_timeout       = 60
  subnets = [
    var.pub_sub_0_id,
    var.pub_sub_1_id,
  ]
  security_groups = [
    var.https_sg_id,
    var.http_redirect_sg_id,
  ]
  tags = {
    Name = "FTRA-${var.env}-ALB"
  }
}

resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.alb.arn
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = var.certificate_arn
  ssl_policy        = "ELBSecurityPolicy-2016-08"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.alb_tg.arn
  }
}

resource "aws_lb_listener" "redirect_https" {
  load_balancer_arn = aws_lb.alb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "redirect"
    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_target_group" "alb_tg" {
  name                 = "FTRA-ALB-TG"
  target_type          = "ip"
  vpc_id               = var.vpc_id
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 300
  health_check {
    path                = "/"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = 200
    port                = "traffic-port"
    protocol            = "HTTP"
  }
  tags = {
    Name = "FTRA_${var.env}_ALB_TG"
  }
  depends_on = [aws_lb.alb]
}

resource "aws_lb_target_group" "api_alb_tg" {
  name                 = "FTRA-API-ALB-TG"
  target_type          = "ip"
  vpc_id               = var.vpc_id
  port                 = 8080
  protocol             = "HTTP"
  deregistration_delay = 300
  health_check {
    path                = "/playground"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = 200
    port                = "traffic-port"
    protocol            = "HTTP"
  }
  tags = {
    Name = "FTRA_${var.env}_API_ALB_TG"
  }
  depends_on = [aws_lb.alb]
}

resource "aws_lb_listener_rule" "alb_listener_rule" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 1
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.alb_tg.arn
  }
  condition {
    path_pattern {
      values = ["/"]
    }
  }
}

resource "aws_lb_listener_rule" "api_alb_listener_rule" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 2
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.api_alb_tg.arn
  }
  condition {
    path_pattern {
      values = ["/api"]
    }
  }
}

resource "aws_lb_listener_rule" "playground_alb_listener_rule" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 3
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.api_alb_tg.arn
  }
  condition {
    path_pattern {
      values = ["/playground"]
    }
  }
}

output "alb_dns_name" {
  value = aws_lb.alb.dns_name
}

output "alb_zone_id" {
  value = aws_lb.alb.zone_id
}

output "alb_terget_group_arn" {
  value = aws_lb_target_group.alb_tg.arn
}

output "api_alb_terget_group_arn" {
  value = aws_lb_target_group.api_alb_tg.arn
}
