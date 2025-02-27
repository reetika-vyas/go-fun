minikube profile minikube

# Clean Kiali (via Helm)
# kubectl delete kiali --all --all-namespaces
# helm uninstall --namespace kiali-operator kiali-operator
# kubectl delete crd kialis.kiali.io

# Remove Istio Objects
istioctl x uninstall -y --purge
# Delete Istio Namespaces
kubectl delete ns istio-system

