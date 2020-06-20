data "template_file" "service_task_definition_json" {
  template = file("../../../modules/ecs/task-definitions/service.json")

  vars = {
    api_image-repository = var.ecr_api_repository_uri
    ui_image-repository  = var.ecr_ui_repository_uri
    dd-agent-key         = var.dd_agent_key
  }
}

data "aws_caller_identity" "self" {}

module "dynamodb" {
  source = "../../../modules/dynamodb"
  env    = var.env
}

module "ecs" {
  source                   = "../../../modules/ecs"
  self                     = data.aws_caller_identity.self.account_id
  env                      = var.env
  alb_terget_group_arn     = data.terraform_remote_state.network.outputs.alb_terget_group_arn
  api_alb_terget_group_arn = data.terraform_remote_state.network.outputs.api_alb_terget_group_arn
  pri_sub_0_id             = data.terraform_remote_state.network.outputs.pri_sub_0_id
  pri_sub_1_id             = data.terraform_remote_state.network.outputs.pri_sub_1_id
  security_group_80_id     = data.terraform_remote_state.network.outputs.pri_app_security_group_id
  security_group_8080_id   = data.terraform_remote_state.network.outputs.pri_app_security_group_8080_id
  container_definitions    = data.template_file.service_task_definition_json.rendered
}

