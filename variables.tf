variable "name" {
  type        = string
  description = "The name for the bucket"
}

variable "owner" {
  type        = map(string)
  description = "Bucket owner's display name and ID. Conflicts with `acl`"
  default     = {}
}

variable "grant" {
  type        = any
  description = "An ACL policy grant. Conflicts with `acl`"
  default     = []
}
