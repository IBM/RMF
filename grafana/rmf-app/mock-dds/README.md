# DDS Mock Server

A mock server that emulates the IBM RMF DDS HTTP API for local development
and testing of the Grafana plugin.

## Usage

```bash
go run . [-port 8803] [-data data]
```

| Flag | Default | Description |
|------|---------|-------------|
| `-port` | `8803` | Port to listen on |
| `-data` | `data` | Directory with JSON snapshot files |

In Grafana, point the datasource at the running server.
