# cka-go

Go application running on Kubernetes.

When developing locally, whenever the application code changes, we re-build the Docker image(s) and refresh the Kubernetes deployment:

* docker build -t go-app:latest . 
* kubectl config use-context kind-first-cluster 
* kind load docker-image go-app:latest --name first-cluster 
* kubectl rollout restart deployment go-deployment

To access the Go app locally:

Find the internal IP of the node where the NodePort go-service is running: 

* kubectl get nodes -o wide
* curl -X POST "http://INTERNAL-IP:30004/set/?key=username&value=JohnDoe"
* curl "http://INTERNAL-IP:30004/get/username" 

Pods are inside Kubernetes nodes, and each node has an internal IP. NodePort services expose a port (30004) on every node in the cluster. Any request to : gets forwarded to the pod inside the node. Alternatively, map a local port to the go-service inside the cluster:

* kubectl port-forward svc/go-service 8000:80 

Intercepts traffic from localhost:8000 on our local machine. Forwards it to port 80 inside the Kubernetes service (go-service). Routes it to the correct pod running the Go app inside the cluster. Effectively, it makes the Go service behave as if it were running directly on localhost:8000