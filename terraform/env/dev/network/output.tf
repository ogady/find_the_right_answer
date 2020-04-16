output "vpc" {
  value = module.vpc.vpc_id ## ここではmoduleのoutputを指定する。
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

output "http_security_group_id" {
  value = module.http_sg.security_group_id
}

output "alb_terget_group_arn" {
  value = module.alb.alb_terget_group_arn
}
