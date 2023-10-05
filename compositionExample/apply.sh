#!/bin/bash
kubectl apply -f xrd.yaml
kubectl apply -f composition.yaml
kubectl apply -f deploy.yaml
kubectl get managed
kubectl get composite
