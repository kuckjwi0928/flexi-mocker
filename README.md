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
- If you format **$variable$** in the response, it will be replaced with a query parameter or path parameter.

```toml
# server config
[server]
port = 8080 # port

# route config
[[routes]]
method = "GET" # http method
path = "/hello" # request path
[routes.response] # response your want to return
hello = "world"
```

### Note
- The body of a POST request is not yet supported for **$variable$** format. (TODO)
