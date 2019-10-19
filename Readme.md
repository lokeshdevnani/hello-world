# Hello World Enterprise Edition

## Running the app locally
```
go build app/hello_world.go
./hello_world
```

## Generate modules
```
go mod tidy
```

## Let's hit the server
```bash
$ curl http://localhost:8080?name=World
Hello, World
```

## Building and running the docker image
```
docker build -t hello-go -f Dockerfile .
docker run -d -p 8080:8080 hello-go
```

## Kubernetes (Minikube)
```
minikube start
```

```
minikube dashboard
```

```
minikube service grafana-service
minikube service hello-world-service
```

## Build & Deploy to k8s
```
sh deploy/deploy.sh
```

## Deploy Observability Loki Stack
```
sh observability/deploy_loki_stack.sh
```

## Reset Everything
```
sh ./reset.sh
```
