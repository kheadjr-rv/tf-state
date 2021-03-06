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

module "sub_foo" {
  source = "./nested_foo"

}
