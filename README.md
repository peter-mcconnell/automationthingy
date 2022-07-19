automationthingy
================

honeycomb.io integration
------------------------

Set the following environment variables before running the program:

```sh
export OTEL_EXPORTER_OTLP_ENDPOINT="grpc://api.honeycomb.io:443"
export OTEL_EXPORTER_OTLP_HEADERS="x-honeycomb-team=o5idz4x55gRW7PaGXqVNBC"
export OTEL_SERVICE_NAME="your-service-name"
```
