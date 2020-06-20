output "vpc" {
  ## ここではmoduleのoutputを指定する。
  value = module.vpc.vpc_id
}

output "pri_sub_0_id" {
  value = module.subnet.pri_sub_0_id
}

output "pri_sub_1_id" {
  value = module.subnet.pri_sub_1_id
}

output "security_group_id" {
  value = module.https_sg.security_group_id
}

output "redirect_security_group_id" {
  value = module.http_redirect_sg.security_group_id
}

output "pri_app_security_group_id" {
  value = module.pri_app_sg_80.security_group_id
}

output "pri_app_security_group_8080_id" {
  value = module.pri_app_sg_8080.security_group_id
}

output "alb_terget_group_arn" {
  value = module.alb.alb_terget_group_arn
}

output "api_alb_terget_group_arn" {
  value = module.alb.api_alb_terget_group_arn
}
