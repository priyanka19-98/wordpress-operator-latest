# wordpress-operator-latest
This repository contains Wordpress Operator built with latest operator-sdk version (v1.6)
## Prerequisites
- golang v1.15+
- set GO111MODULE="on"
- [Install the operator-sdk (v1+)](https://docs.openshift.com/container-platform/4.7/operators/operator_sdk/osdk-installing-cli.html)
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [kubectl client](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Trying the Operator

```
git clone https://github.com/priyanka19-98/wordpress-operator-latest.git
cd wordpress-operator-latest
```
We would be trying out the operator locally. By locally we mean that we want to run the operatot logic binary without actually building an image and pushing it to a container registry. Running the operator locally helps in day to day development. 
You can have a [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) single node local cluster to play with the operator.

`make install run` 

From the other terminal window, create the Wordpress CR: 
```
kubectl create -f config/samples/wordpress_v1_wordpress.yaml
``` 

Your `config/samples/wordpress_v1_wordpress.yaml` must look something like:
```
apiVersion: wordpress.example.com/v1
kind: Wordpress
metadata:
  name: mysite
spec:
  sqlrootpassword: "abc"
 ```

You should be able to see the output of `make install run` similar to:
```
/home/pjiandan/go/src/wordpress-operator/bin/controller-gen "crd:trivialVersions=true,preserveUnknownFields=false" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/home/pjiandan/go/src/wordpress-operator/bin/kustomize build config/crd | kubectl apply -f -
customresourcedefinition.apiextensions.k8s.io/wordpresses.wordpress.example.com created
/home/pjiandan/go/src/wordpress-operator/bin/controller-gen "crd:trivialVersions=true,preserveUnknownFields=false" rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases
/home/pjiandan/go/src/wordpress-operator/bin/controller-gen object:headerFile="hack/boilerplate.go.txt" paths="./..."
go fmt ./...
go vet ./...
go run ./main.go
2021-07-05T14:15:59.495+0530    INFO    controller-runtime.metrics      metrics server is starting to listen    {"addr": ":8080"}
2021-07-05T14:15:59.496+0530    INFO    setup   starting manager
2021-07-05T14:15:59.497+0530    INFO    controller-runtime.manager      starting metrics server {"path": "/metrics"}
2021-07-05T14:15:59.497+0530    INFO    controller-runtime.manager.controller.wordpress Starting EventSource    {"reconciler group": "wordpress.example.com", "reconciler kind": "Wordpress", "source": "kind source: /, Kind="}
2021-07-05T14:15:59.598+0530    INFO    controller-runtime.manager.controller.wordpress Starting EventSource    {"reconciler group": "wordpress.example.com", "reconciler kind": "Wordpress", "source": "kind source: /, Kind="}
2021-07-05T14:15:59.698+0530    INFO    controller-runtime.manager.controller.wordpress Starting EventSource    {"reconciler group": "wordpress.example.com", "reconciler kind": "Wordpress", "source": "kind source: /, Kind="}
2021-07-05T14:15:59.799+0530    INFO    controller-runtime.manager.controller.wordpress Starting EventSource    {"reconciler group": "wordpress.example.com", "reconciler kind": "Wordpress", "source": "kind source: /, Kind="}
2021-07-05T14:15:59.900+0530    INFO    controller-runtime.manager.controller.wordpress Starting Controller     {"reconciler group": "wordpress.example.com", "reconciler kind": "Wordpress"}
2021-07-05T14:15:59.900+0530    INFO    controller-runtime.manager.controller.wordpress Starting workers        {"reconciler group": "wordpress.example.com", "reconciler kind": "Wordpress", "worker count": 1}
2021-07-05T14:15:59.900+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:15:59.900+0530    INFO    controllers.Wordpress   Creating a new PVC      {"PVC.Namespace": "default", "PVC.Name": "mysql-pv-claim"}
2021-07-05T14:15:59.947+0530    INFO    controllers.Wordpress   Creating a new Deployment  {"Deployment.Namespace": "default", "Deployment.Name": "wordpress-mysql"}
2021-07-05T14:15:59.969+0530    INFO    controllers.Wordpress   Creating a new Service  {"Service.Namespace": "default", "Service.Name": "wordpress-mysql"}
2021-07-05T14:15:59.977+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:15:59.978+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:15:59.978+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:15:59.998+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:15:59.998+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:16:00.044+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:16:00.044+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:16:00.058+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:16:00.059+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:17:09.986+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:17:09.987+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:17:24.988+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:17:24.988+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:18:09.993+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:18:09.993+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:18:14.993+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:18:14.994+0530    INFO    controllers.Wordpress   MySQL isn't running, waiting for 5s
2021-07-05T14:19:10.278+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:19:10.278+0530    INFO    controllers.Wordpress   Creating a new PVC      {"PVC.Namespace": "default", "PVC.Name": "wp-pv-claim"}
2021-07-05T14:19:10.292+0530    INFO    controllers.Wordpress   Creating a new Deployment  {"Deployment.Namespace": "default", "Deployment.Name": "wordpress"}
2021-07-05T14:19:10.309+0530    INFO    controllers.Wordpress   Creating a new Service  {"Service.Namespace": "default", "Service.Name": "wordpress"}
2021-07-05T14:19:10.341+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:19:10.355+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:19:10.361+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:19:10.371+0530    INFO    controllers.Wordpress   Reconciling Wordpress
2021-07-05T14:19:15.001+0530    INFO    controllers.Wordpress   Reconciling Wordpress
``` 

From the different terminal window, check for the resources created:

```
[pjiandan@localhost wordpress-operator]$ kubectl get all
NAME                                   READY   STATUS    RESTARTS   AGE
pod/wordpress-7446b985d9-vc7w2         1/1     Running   0          128m
pod/wordpress-mysql-5cd8987844-4p6jp   1/1     Running   0          128m

NAME                      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
service/kubernetes        ClusterIP   10.96.0.1       <none>        443/TCP        3h58m
service/wordpress         NodePort    10.110.12.200   <none>        80:32324/TCP   128m
service/wordpress-mysql   ClusterIP   None            <none>        3306/TCP       128m

NAME                              READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/wordpress         1/1     1            1           128m
deployment.apps/wordpress-mysql   1/1     1            1           128m

NAME                                         DESIRED   CURRENT   READY   AGE
replicaset.apps/wordpress-7446b985d9         1         1         1       128m
replicaset.apps/wordpress-mysql-5cd8987844   1         1         1       128m
```

Next, run the following command to return the IP address for the WordPress service:
```
[pjiandan@localhost wordpress-operator]$ minikube service wordpress --url
http://192.168.99.143:32324
```
Open the url in your browser and you should see: 
![alt text](https://raw.githubusercontent.com/kubernetes/examples/master/mysql-wordpress-pd/WordPress.png)


# Questions
Please feel free to open up an issue.


