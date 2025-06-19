# go-rest-templ

A simple template for a REST API server in Go.

## Run

Run the server with `SERVER_CONFIG_PATH=local_config.yaml go run cmd/server/main.go`. An example config file is given below:

```yaml
port: 9999

cat-facts:
  enabled: true
  url: "https://catfact.ninja/fact"
```

I come back to this from time to time and update it, currently using the following:

- config: YAML file via env var
- http: chi+huma
- docs: huma
- logging: zerolog
- metrics: prometheus
- CI/CD: GHA (TODO)
