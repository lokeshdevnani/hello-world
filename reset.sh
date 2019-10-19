# delete all kubernetes resource in cluster
kubectl config use-context minikube
kubectl -n default delete pod,svc,deploy,ds,sts --all
