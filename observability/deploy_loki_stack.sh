#!/bin/sh

kubectl config use-context minikube

cd `dirname $0`

for file in $(ls loki/*.yaml); do
	kubectl apply -f $file
done

for file in $(ls promtail/*.yaml); do
	kubectl apply -f $file
done

for file in $(ls prometheus/*.yaml); do
	kubectl apply -f $file
done

for file in $(ls grafana/*.yaml); do
	kubectl apply -f $file
done
