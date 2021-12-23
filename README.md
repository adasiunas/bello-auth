# Bello auth app
Simple go app to showcase kubernetes deployments on minikube.

## Prerequisites for your local environment in case you'd like to test it out:
- Minikube (installation guide: https://minikube.sigs.k8s.io/docs/start/)
- `kubectl` (installation guide: https://kubernetes.io/docs/tasks/tools/)
- Docker (installation guide: https://www.docker.com/products/docker-desktop)
- Go (installation guide: https://go.dev/doc/install)
- This service API is auto-generated with `goswagger` binary using `swagger.json` file, thus for local development `goswagger` installation will be needed (guide & docs: https://goswagger.io/install.html **note: I've used v0.19.0**)

## Starting service locally:
- Make sure that all installations listed above are in place.
- In case it is your first time, you will need to setup a mysql db instance & create db schema used in this service:
	- Download mysql docker image from Docker hub & start a container:
		```
		$ docker run --name bello-auth-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag
	- Connect to the mysql container & create DB schema needed for this service:
		- List running docker containers:
			```
			$ docker container ls
		- Find mysql container sha and paste it into command listed below:
			```
			$ docker exec -it <bello-mysql-container-sha> bash
		- You should be connected to the container instance via bash, connect to the mysql server using root user (enter the pass you've used when starting a container in the previous step)
			```
			$ root@<container-sha>:/# mysql -u root -p
		- Copy & paste content from `db/migration/mysql_init.sql` file into the terminal.
		- DB schema should be successfully created, in order to confirm that you can use following commands inside mysql client:
		```
		$ mysql> USE bello_auth; # This should change current context database
			Database changed
		$ mysql> SHOW TABLES;
			+----------------------+
			| Tables_in_bello_auth |
			+----------------------+
			| user                 |
			+----------------------+

- Once your local database is ready, you are good to go & start this service locally:
	- Inside your terminal you can `source` local environment variables (those can be found in `local.env` file):
		```
		$ source local.env
	- Use `make run` command to run the service (or use `go run cmd/bello-app-auth-server/main.go` in case Makefile is not recognized)
		```
		$ make run
		go run cmd/bello-app-auth-server/main.go
		..... Serving bello app auth at http://127.0.0.1:5000
	- Application should be started on `localhost:5000` (to confirm that you can check app swagger docs on: `localhost:5000/docs` via your browser)


## Starting service on Minikube:
In short minikube is a _small_ kubernetes cluster for your workstation which can be used as a sandbox for mitigating your development environment. For more details on that you can check out their documentation.

- Start minikube cluster:
	```
	$ minikube start
- Via `kubectl` you should be able to see `minikube` cluster:
	```
	$ kubectl config get-contexts
	CURRENT   NAME       CLUSTER    AUTHINFO                               NAMESPACE
	*         minikube   minikube   minikube                               default
- Make sure that asterisk (`*`) symbol in `CURRENT` column is next to the `minikube` cluster.
- We will be working on a default namespace inside minikube cluster when creating deployments for this app.
- All the deployment configuration files for this app are inside `k8s-config/` folder.
- Apply configuration via `kubectl` command in the following order:
	- Creating mysql database service & secrets:
		```
		$ kubectl apply -f k8s-config/bello-mysql/bello-mysql-secret.yaml
		$ kubrctl apply -f k8s-config/bello-mysql/bello-mysql.yaml
	- To verify if `mysql` service is up and running you can grep pods inside our minikube cluster:
		```
		$ kubectl get pods
		NAME                          READY   STATUS    RESTARTS       AGE
		bello-mysql-744ff984c-cpw24   1/1     Running   1 (164m ago)   1h
	- Creating bello-app deployment, configmap & load balancer:
		```
		$ kubectl apply -f k8s-config/bello-app/bello-app-configmap.yaml
		$ kubectl apply -f k8s-config/bello-app/bello-app.yaml
		$ kubectl apply -f k8s-config/bello-app/bello-app-load-balancer.yaml
	- To verify if service is up and running you can grep services in our cluster:
		```
		$ kubectl get services
		NAME                     TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
		bello-app-loadbalancer   LoadBalancer   10.107.200.22   <pending>     8000:30000/TCP   1d
		bello-mysql-service      ClusterIP      10.111.5.249    <none>        3306/TCP         1d
		kubernetes               ClusterIP      10.96.0.1       <none>        443/TCP          1d
		```
		```
		$ kubectl get pods
		NAME                          READY   STATUS    RESTARTS       AGE
		bello-app-576fbf9b78-kzrq2    1/1     Running   3 (175m ago)   1d
		bello-mysql-744ff984c-cpw24   1/1     Running   1 (176m ago)   1d
	- For accessing service from your browser, you can tell minikube to serve `bello-app-loadbalancer` service with an external URL (in another Terminal session):
		```
		$ minikube service bello-app-loadbalancer
		|-----------|------------------------|-------------|---------------------------|
		| NAMESPACE |          NAME          | TARGET PORT |            URL            |
		|-----------|------------------------|-------------|---------------------------|
		| default   | bello-app-loadbalancer |        8000 | http://<some-IP>:30000 |
		|-----------|------------------------|-------------|---------------------------|
		🎉  Opening service default/bello-app-loadbalancer in default browser...
	- To verify if it works, you can visit `http://<some-IP>:30000/docs` page and swagger docs should be rendered.