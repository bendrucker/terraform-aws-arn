output "arn" {
  value = join(":", [
    "arn",
    var.partition,
    var.service,
    var.region,
    var.account_id,
    var.resource_id
  ])

  description = "The ARN in string form"
}
