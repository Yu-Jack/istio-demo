## Introduction

This is demo project for istio.  

When you curl `curl /product?version=v2`, you always get `this is v2 review` message. Otherwise, you only get `this is v1 review`.

## Bootstrap

1. `make up`
2. `kubectl port-forward deployment/product-v1 8080:8080`
3. Try to curl `http://localhost:8080/product?version=v1` or `http://localhost:8080/product?version=v2`

p.s. The author uses minikube.