#!/bin/bash
kubectl delete -f deploy.yaml
kubectl delete -f composition.yaml
kubectl delete -f xrd.yaml
kubectl get managed
kubectl get composite
