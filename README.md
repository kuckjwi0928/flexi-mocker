# Flexi Mocker

## Getting Started

### Build locally

```shell
$ make build
```

### Run on local
```shell
$ make run
```

### Configuration
- You can configure mocking by defining a `config.toml` file.
- If you format **$query.variableKey** in the response, it will be replaced with a query parameter.
- If you format **$path.variableKey** in the response, it will be replaced with a path parameter.
- If you format **$header.variableKey** in the response, it will be replaced with a header parameter.
- If you format **$body.variableKey** in the response, it will be replaced with a body parameter.

```toml
# server config
[server]
port = 8080 # port

# route config
[[routes]]
method = "GET" # http method
path = "/hello/:path-variable" # request path
[routes.response] # response your want to return
hello = "world"
path = "$path-variable$"
```
