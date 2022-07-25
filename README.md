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

github app
----------

In the `.automationthingy.yaml` file you can see that the api is configured
with some auth: github: values. These values expect a secret to be available in
vault. After creating your Github app for oauth2 authentication please record
your secret in vault with:

```sh
kubectl exec -n vault -ti vault-0 -- /bin/sh -c "vault kv put kv-v1/keys/githubapp secret='<my github app secret>'"
```
