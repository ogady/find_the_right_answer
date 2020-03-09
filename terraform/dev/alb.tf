resource "aws_lb" "alb" {
  name               = "FTRA-ALB"
  load_balancer_type = "application"
  internal           = false
  idle_timeout       = 60
  subnets = [
    aws_subnet.pub_subnet_0.id,
    aws_subnet.pub_subnet_1.id,
  ]
  security_groups = [
    module.http_sg.security_group_id,
    module.https_sg.security_group_id,
    module.http_redirect_sg.security_group_id
  ]

}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.alb.arn
  port              = "80"
  protocol          = "HTTP"
  default_action {
    type = "fixed-response"
    fixed_response {
      content_type = "text/plain"
      message_body = "これはHTTPです"
      status_code  = "200"
    }
  }
}
resource "aws_lb_target_group" "alb_tg" {
  name                 = "FTRA-ALB-TG"
  target_type          = "ip"
  vpc_id               = aws_vpc.vpc.id
  port                 = 8080
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
    Name = "FTRA_ALB_TG"
  }
  depends_on = [aws_lb.alb]
}

resource "aws_lb_listener_rule" "alb_listener_rule" {
  listener_arn = aws_lb_listener.http.arn
  priority     = 100
  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.alb_tg.arn
  }
  condition {
    field  = "path-pattern"
    values = ["/*"]
  }

}

