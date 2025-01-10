To run a pod - ```sudo kubectl run nginx --image=nginx```

To get the pods - ```sudo kubectl get pods```

To get the pods in all namespaces
```sudo kubectl get pods --all-namespaces```

To set the context to a specific namespace
```sudo kubectl config set-context --current --namespace=dev```

To get-context
```sudo kubectl config get-contexts```

To set-context
```sudo kubectl config use-context dev-context```

