# superheroe-gokit-api

a simple go-kit api that provides information about superheroes

# instalation with docker and docker compose

- make sure you have docker and docker-compose installed

- in the root of the proyect, execute docker-compose up and compose will carry all the hard work to activate the app

# instalation with kubernetes

- make sure to have installed minikube or any other local kubernetes cluster

- in the root of the proyect run kubectl -f apply k8s/

- then describe the superheroe-gokit-api-svc and use that ip with 8080 port

- if you are using minikube, remember to apply a minikube tunnel in order to communicate with the cluster

# run without docker or docker compose

- if you have Go installed, you can execute the main.go file inside the src folder, or run the Make build command inside the Makefile and then execute the build allocated inside build/bin folder