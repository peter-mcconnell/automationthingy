.PHONY: build run run_and_update k3s vaultdev pprof
build:
	go build -o automationthingy

run: build
	./automationthingy

run_and_update:
	@if ! command -v gow > /dev/null; then echo "please install 'gow'. go install github.com/mitranim/gow@latest" && exit 1; fi
	gow run .

pprof:
	@if [ ! -f automationthingy.pprof ]; then echo "automationthingy.pprof not found. run binary with -pprof to create it" && exit 1; fi
	go tool pprof -http=":9080" ./automationthingy automationthingy.pprof

k3s:
	@if ! command -v k3s > /dev/null; then echo "please install 'k3s'" && exit 1; fi
	k3s server --docker

vaultdev:
	@if ! command -v helm > /dev/null; then echo "please install 'helm'" && exit 1; fi
	@if ! command -v kubectl > /dev/null; then echo "please install 'kubectl'" && exit 1; fi
	@if ! kubectl get po -n vault vault-0 > /dev/null; then \
		echo " ~ installing vault" && \
		helm repo add hashicorp https://helm.releases.hashicorp.com && \
		helm install vault hashicorp/vault --create-namespace --namespace vault --set "server.dev.enabled=true" --wait; \
	else \
		echo " ~ already running vault. nothing to do"; \
	fi
	@if ! kubectl exec -n vault -it vault-0 -- /bin/sh -c 'vault secrets list | grep -q kv-v1/'; then \
		echo " ~ vault: enabling secrets engine at kv-v1" && \
		kubectl exec -n vault -it vault-0 -- /bin/sh -c 'vault secrets enable -path="kv-v1" -description="automationthingy kv" kv'; \
	else \
		echo " ~ vault: secrets engine already enabled at kv-v1"; \
	fi
	@if ! kubectl exec -n vault -it vault-0 -- /bin/sh -c 'vault secrets list | grep -q key-automationthingy'; then \
		if [ ! -f "$(PWD)/key-automationthingy" ]; then \
			echo " ~ creating private key"; \
			ssh-keygen -t rsa -b 4096 -C "automationthingy" -f "$(PWD)/key-automationthingy" -N ''; \
		else \
			echo " ~ not creating private key - $(PWD)/key-automationthingy already exists"; \
		fi; \
		echo " ~ vault: creating secret kv-v1/keys/key-automationthingy"; \
		kubectl exec -n vault -it vault-0 -- /bin/sh -c "vault kv put kv-v1/keys/key-automationthingy cert='$$(cat key-automationthingy)'"; \
	else \
		echo " ~ vault: secret kv-v1/keys/key-automationthingy already exists"; \
	fi
