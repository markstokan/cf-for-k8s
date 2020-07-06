#!/bin/bash
allowlist='cf\-blobstore\-minio|kpack\-webhook'
output_file=/tmp/output.txt
touch $output_file 
for namespace in $(kubectl get namespaces -o jsonpath='{.items..metadata.name}') ; do
  if [ "$namespace" != "istio-system" ] && [ "$namespace" != "kube-system" ]; then
    for pod in $(kubectl get pod -n ${namespace} -o jsonpath='{.items..metadata.name}') ; do
	istioctl authn tls-check ${pod}.${namespace} >> $output_file
    done
  fi
done

if cat  $output_file | grep -i permissive | sort | uniq | awk '{print $1}' | grep -Ev "$allowlist"; then
  echo "found unknown permissive connections"
  exit 1
fi
  echo "verified no new permissive connections"


