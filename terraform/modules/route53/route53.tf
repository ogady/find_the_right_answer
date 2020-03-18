variable "alb_dns_name" {}
variable "alb_zone_id" {}

data "aws_route53_zone" "route53_zone" {
  name = "ogady.net"
}

resource "aws_route53_record" "route53_record" {
  zone_id = data.aws_route53_zone.route53_zone.zone_id
  name    = data.aws_route53_zone.route53_zone.name
  type    = "A"

  alias {
    name                   = var.alb_dns_name
    zone_id                = var.alb_zone_id
    evaluate_target_health = true
  }

}

resource "aws_acm_certificate" "acm" {
  domain_name               = aws_route53_record.route53_record.name
  subject_alternative_names = []
  validation_method         = "DNS"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "sample_certificate" {
  name    = aws_acm_certificate.acm.domain_validation_options[0].resource_record_name
  type    = aws_acm_certificate.acm.domain_validation_options[0].resource_record_type
  records = [aws_acm_certificate.acm.domain_validation_options[0].resource_record_value]
  zone_id = data.aws_route53_zone.route53_zone.id
  ttl     = 60
}

resource "aws_acm_certificate_validation" "sample_valid" {
  certificate_arn         = aws_acm_certificate.acm.arn
  validation_record_fqdns = [aws_route53_record.sample_certificate.fqdn]
}

output "domain_name" {
  value = aws_route53_record.route53_record.name
}

output "certificate_arn" {
  value = aws_acm_certificate.acm.arn
}


