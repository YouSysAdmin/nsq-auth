// Alpine version
variable "GOLANG_VERSION" {
  default = "latest"
}

variable "APP_VERSION" {
  default = "local"
}

target "args" {
  args = {
    GOLANG_VERSION = GOLANG_VERSION
    APP_VERSION = APP_VERSION
  }
}

target "platforms" {
  platforms = [
    "linux/amd64",
    "linux/arm64",
    "linux/arm/v7"
  ]
}

// Special target: https://github.com/docker/metadata-action#bake-definition
target "docker-metadata-action" {
  tags = ["nsq-auth:${APP_VERSION}"]
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["args", "docker-metadata-action"]
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
}

target "image-all" {
  inherits = ["platforms", "image"]
  #tags=["ghcr.io/yousysadmin/nsq-auth:${APP_VERSION}"]
}
