ts=`date +%s`

cd `dirname $0`
image_path=lokeshdevnani/hello-go

echo "Building Docker image..."
docker build -t hello-go -f ../Dockerfile ..
docker tag hello-go ${image_path}:${ts}
docker push ${image_path}:${ts}
docker tag hello-go ${image_path}:latest
# gcloud container images add-tag --quiet ${image_path}:${ts} ${image_path}:latest

echo "Deploying to k8s"
kubectl apply -f hello-world-deployment.yaml
kubectl apply -f hello-world-service.yaml
kubectl set image deployment hello-world hello-world=${image_path}:${ts} --record
kubectl rollout status deployment hello-world