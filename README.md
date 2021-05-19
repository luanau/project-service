# Setting up a Go service on kubernetes using a dapr redis state store
This is just an example of setting up a Go service on a kubernetes with DAPR. The service is using a dapr redis state component to save data. The deployment environment is kubernetes that comes with Docker desktop for windows. The Go code is very simplistic as I am just starting to play with Go.
## Setup DAPR on your kubernetes (Docker desktop for windows)
1. Install dapr to k8s with helm as shown at [this page](https://docs.dapr.io/operations/hosting/kubernetes/kubernetes-deploy/) 
2. Deploy redis to k8s with helm
    - Instructions is at [this page](https://docs.dapr.io/reference/components-reference/supported-state-stores/setup-redis/#configuration)
## Application setup
- Add dapr state component
- Build docker image of your app/api
- Deploy your api on k8s (with dapr annotation for your pod)
- Test your app/api

1. Deploy a dapr component to use the above redis as a store (using deploy/redis.yaml)
    - This yaml is compatible with above redis deployment
    ```
    kubectl apply -f deploy/redis.yaml
    ```
2. Build an image of your program,
    ```
    docker build -t project-service .
    ```
3. Deploy your program (deployment, service, ingress) to k8s
    ```
    kubectl apply -f deploy/project-service.yaml
    ```
4. Test your api at 'alpaca.example.com/project-service'
This is the uri you set up with your Ingress in file deploy/project-service.yaml
    - Just hello world and extra url params
    ```
    curl http://alpaca.example.com/project-service
    curl http://alpaca.example.com/project-service/query?name=Fred&age=93
    ```
    - Post a project json to the api. It will return the object saved.
    ```
    curl -X POST -H "Content-Type: application/json" -d @project.json http://alpaca.example.com/project-service
    ```

## Checking the redis store
```
kubectl get pods

ingress-nginx-controller-57cb5bf694-2n68f   1/1     Running   3          39h  
project-service-78f99c65f6-qq94c            2/2     Running   0          79m  
redis-master-0                              1/1     Running   6          2d16h
redis-replicas-0                            1/1     Running   7          2d16h
redis-replicas-1                            1/1     Running   6          2d16h
redis-replicas-2                            1/1     Running   6          2d16h
```
### Start redis-cli on pod redis-master-0 
```
k exec -it redis-master-0 -- redis-cli
```
### Issue redis-cli commands to interact
```
auth <password>
keys *
hgetall <key>
```