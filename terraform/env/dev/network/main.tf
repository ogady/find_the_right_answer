module "vpc" {
  source = "../../../modules/vpc"
  env    = var.env
}

module "subnet" {
  source   = "../../../modules/subnet"
  env      = var.env
  vpc_id   = module.vpc.vpc_id
  vpc_cidr = module.vpc.vpc_cidr
}

module "alb" {
  source              = "../../../modules/alb"
  env                 = var.env
  vpc_id              = module.vpc.vpc_id
  pub_sub_0_id        = module.subnet.pub_sub_0_id
  pub_sub_1_id        = module.subnet.pub_sub_1_id
  pri_sub_0_id        = module.subnet.pri_sub_0_id
  pri_sub_1_id        = module.subnet.pri_sub_1_id
  http_sg_id          = module.http_sg.security_group_id
  https_sg_id         = module.https_sg.security_group_id
  http_redirect_sg_id = module.http_redirect_sg.security_group_id
  certificate_arn     = module.route53.certificate_arn
}

module "http_sg" {
  source      = "../../../modules/security_group"
  name        = "http_sg"
  vpc_id      = module.vpc.vpc_id
  port        = 80
  cidr_blocks = ["0.0.0.0/0"]
}


module "https_sg" {
  source      = "../../../modules/security_group"
  name        = "https_sg"
  vpc_id      = module.vpc.vpc_id
  port        = 443
  cidr_blocks = ["0.0.0.0/0"]
}

module "http_redirect_sg" {
  source      = "../../../modules/security_group"
  name        = "http_redirect_sg"
  vpc_id      = module.vpc.vpc_id
  port        = 8080
  cidr_blocks = ["0.0.0.0/0"]
}

module "route53" {
  source       = "../../../modules/route53"
  alb_dns_name = module.alb.alb_dns_name
  alb_zone_id  = module.alb.alb_zone_id
}

