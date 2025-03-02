# Deploy & run application via Kubernetes using Minikube cluster
Reference: https://kubernetes.io/docs/tutorials/hello-minikube/

**Install minikube to run it as a container (For Mac M1)**
    eval "$(/opt/homebrew/bin/brew shellenv)"
    arch -arm64 brew install minikube
    arch -arm64 minikube start
    arch -arm64 minikube dashboard

Permanent fix: 
echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> ~/.zprofile
source ~/.zprofile
minikube start
minikube dashboard

# Kubernetes
kubectl config use-context minikube
kubectl create ns todo-app
1. Create a configmap to store postgres database environment variables
    kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/postgres_configmap.yaml --namespace todo-app
    kubectl get configmaps -n todo-app
    kubectl describe cm/postgres-credentials -n todo-app
2. Create PersistentVolumeClaim for postgres
   kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/postgres_pvc.yml --namespace todo-app
   kubectl get pvc -n todo-app
   kubectl describe pvc/postgres-pvc -n todo-app
3. Create a service as nodeport for postgres
   kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/postgres_service.yaml -n todo-app
   kubectl get svc -n todo-app
   kubectl describe svc/postgres-service -n todo-app
4. Create a deployment for postgres
   kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/postgres_deployment.yaml -n todo-app
   kubectl get deploy -n todo-app
   kubectl describe deploy/postgres-deploy -n todo-app
   kubectl get pods -n todo-app
5. Create backend k8s service for todoapp
   kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/todo_backend_svc.yaml -n todo-app
   kubectl get svc -n todo-app
   kubectl describe svc/todo-backend-svc -n todo-app
5. Create backend deployment for todoapp
   kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/todo_backend_deploy.yaml -n todo-app
   kubectl get deploy -n todo-app
6. Create ingress for todoapp
   kubectl apply -f ./GolandProjects/TodoListApp/kubernetes/todo_backend_ingress.yaml -n todo-app
   kubectl get ingress -n todo-app

**To connect to postgresDB in minikube**
kubectl port-forward svc/postgres-service 5432:5432
Now, PostgreSQL is available at localhost:5432, connect with Dbeaver.

**To get url to hit APIs using postman/curl**
Expose the service & test the API by command: 
minikube service todo-backend-svc --url
It will give you a URL like: http://<minikube-ip>:<port>
ex. http://192.168.49.2:30001

NOTE: Ingress is not working here. Need to fix it.
==================================================================================================

**PVC (Persistent Volume Claim)**
What Does "Mounts the PVC for Data Storage" Mean?
In Kubernetes, when we say that a PVC (PersistentVolumeClaim) is mounted for data storage, it means:
A Persistent Volume (PV) is allocated: A storage space is reserved for the pod.
The Pod uses the PVC to store its data: Instead of storing data inside the pod (which would be lost if the pod restarts), it is saved on a persistent volume.
The database can restart without losing data: Since the PVC remains even if the pod is recreated, PostgreSQL retains all the stored records.

🛠 What Happens in the Background?
The PVC (Persistent Volume Claim) requests storage from the Kubernetes cluster.
A Persistent Volume (PV) is allocated (if available).
The PostgreSQL pod mounts this storage at /var/lib/postgresql/data.
Now, PostgreSQL stores its database files there instead of inside the pod itself.
Even if the pod is deleted and restarted, the database files remain intact.

🎯 Why Is This Important?
✔ Ensures data persistence (No data loss when the pod restarts)
✔ Decouples storage from the pod (Storage is independent of the pod lifecycle)
✔ Supports scaling (Data remains accessible even if the deployment changes)
