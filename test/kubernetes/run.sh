#!/bin/bash

REPO_HOME=$(git rev-parse --show-toplevel)

# clean up
kubectl delete namespace test --ignore-not-found=true --wait=true
kubectl create namespace test


# deploy
kubectl apply -k ${REPO_HOME}/test/kubernetes/test
kubectl set image -n test deployment/example example=ghcr.io/rakutentech/passenger-go-exporter/passenger-app:$PASSENGER_VERSION
kubectl scale -n test deploy/example --replicas=1
kubectl -n test rollout status deploy/example

POD_IP=$(kubectl get pod -n test -o=jsonpath='{.items[0].status.podIP}')

# test
kubectl run curl -it --rm -n test --restart=Never --image=curlimages/curl:latest -- http://${POD_IP}:9768/metrics > metrics.txt
metrics=(
  "passenger_go_process_count"
  "passenger_go_process_processed"
  "passenger_go_process_real_memory_bytes"
  "passenger_go_wait_list_size"
)
for metric in "${metrics[@]}"
do
  echo "check ${metric} metrics."
  grep -E "^${metric}.*" metrics.txt > /dev/null
  if [ $? -ne 0 ]; then
    echo "not found ${metric}"
    exit 1;
  fi
done
echo "all metrics exists."