#!/bin/bash

set -euo pipefail

function errxit() {
  echo $1 > /dev/stderr
  exit $2
}

function get-docker-password() {
  CTX=`kubectl config current-context`

  case $CTX in
    juvo-ilab-prod) NS=staging-kn ;;
    juvo-dev) NS=dev-kn ;;
    *) errxit "Invalid context $CTX" 1 ;;
  esac

  kubectl get secret dockerhub -n$NS -ojsonpath='{.data.password}' |
  base64 -d
}

get-docker-password |
docker login -u jurisfuturasb --password-stdin

for i in build run; do
  artifact="jurisfutura/juvo-cnb-$i"
  docker buildx build . -t $artifact --target $i
  docker push $artifact
done

