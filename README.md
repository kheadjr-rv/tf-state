# TF State

Quick and simple demonstration how to visualize modules within state.

Simply pipe the state to the application and it will output for the purpose of this demo mermaid.js formatted output.

Using the sample terraform code in the repo.

```
cat terraform/terraform.tfstate | go run main.go
```

Graph output
```mermaid
graph LR
  module.foo-->module.foo.module.sub_foo(module.sub_foo)
  root-->module.bar(module.bar)
  module.foo.module.sub_foo-->module.foo.module.sub_foo.module.third_level(module.third_level)
  root-->module.foo(module.foo)
```

Visualized with mermaid.js [live editor](https://mermaid-js.github.io/mermaid-live-editor/#)

![modules](diagram.png)