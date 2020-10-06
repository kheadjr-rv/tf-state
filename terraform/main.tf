provider "random" {

}

resource "random_id" "server" {
  count       = 2
  byte_length = 8
}

resource "random_id" "task" {
  count       = 0
  byte_length = 8
}

resource "random_id" "service" {
  for_each    = toset(["Todd", "James", "Alice", "Dottie"])
  byte_length = 8
}

resource "random_id" "single" {
  byte_length = 8
}

module "bar" {
  source = "./modules/bar"

}

module "foo" {
  source = "./modules/foo"
}
