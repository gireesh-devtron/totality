This is a monorepo of two repos.
1) server : this repo has the actual grpc server for users microservice
2) client : this repo is a tester client which connects with the above server and tests the functionalities of the server.


# How to setup the server?

### Method1 : using docker
-> assuming you have docker setup 
#### setup user service grpc server : 
1) download the code into your local system.
2) cd to the server directory.
3) run ```docker build . --tag={your-custom-image-tag}```
4) spin up a container using the above built image using ```docker run {image-tag-that-was-created-above}```
5) now the user service grpc server is up and running, by default the server listens on port``3000``. If you like to configure this port you can pass ``port`` env variable with custom port number.

#### setup test client : 

## Method2 : run the server locally on your machine.
-> assuming you have go installed locally.
1) download the code into your local system.
2) open the terminal in ``totality`` directory.
3) run ``cd server`` and run ``go mod tidy`` and run ``go run .``
4) that's it user service grpc server is running on port ```3000``` and you can ypu the service at ```localhost:3000```.

-> how to download proto file:
1) if you like to interact with user service, you can use the client code that was there in the totality directory. this client already has the code to connect with user service. by default client assumes the user service is runnig on port ``3000``. u can configure the port if user service is being run at different port.
2) if you like to create to own service that want to communicate with the user service server, you have to get proto files from the server.
   1) please get the proto files from ``github.com/gireesh-devtron/totality/server/grpc`` go package.
   2) now u can import into your own service and pass the grpc connection object to ``NewUserServiceClient`` function from above package. this function return the ``userService`` u can use this service to communicate with the api's. 
