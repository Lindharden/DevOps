# Lecture 2: Packaging applications, Containerization with Docker
February 07 and 10, 2023

## Step 1: Refactor ITU-MiniTwit to another language and technology of your choice

First of all we ran our test suite to see whether the tests pass or not. Here we have some tests that fail. To fix all of these tests, we had to change the string encoding. 

For our refactoring we decided to use the language GO. We choose this language as it's much faster than Python, and because we wanted to learn it. We installed an GO SQL package on the system, and imported it into our application. This allows us to query the database using `sql.open()`. Then we wanted to update our test suite to support our GO application. We had to update the flask HTML templates to work with GO. Here we had to change the expression declarations from using the following syntax: `{% expression %}` to `{{expression}}`. We then changed the templates to use some new controllers and routes we created for handling different tasks. We created controllers, routes and middleware as described in this blog by Jes Fink-Jensen: <https://betterprogramming.pub/how-to-create-a-simple-web-login-using-gin-for-golang-9ac46a5b0f89>.


## Step 2: Containerize ITU-MiniTwit with Docker

To containerize the application we created an Dockerfile like such:

```
# syntax=docker/dockerfile:1

# Fetch GO
FROM golang:latest

# Create sub directory
WORKDIR /app

# Copy everything from project into container
COPY . ./

# Download all dependencies from go mod
RUN go mod download

# Build binary
RUN go build -o /minitwit

EXPOSE 8080

CMD [ "/minitwit" ]
```
We then build the docker container using the following command: \
`docker build -f Dockerfile -t lind/minitwit .`

Using `docker images` shows us that the container has been created.
```
lind@pop-os:~/Desktop/DevOps/DevOps$ docker images
REPOSITORY      TAG       IMAGE ID       CREATED          SIZE
lind/minitwit   latest    695f1caaf585   32 seconds ago   1.07GB
...
```

We can now run our application using `docker run -d -p 8080:8080 --name minitwit lind/minitwit`. We can now show all running applications in docker using `docker ps -a`, which yields:
```
lind@pop-os:~/Desktop/DevOps/DevOps$ docker ps
CONTAINER ID   IMAGE           COMMAND       CREATED         STATUS         PORTS                                       NAMES
4994d39f2afc   lind/minitwit   "/minitwit"   2 minutes ago   Up 2 minutes   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp   minitwit
```

To ease the process of running the application with the correct ports, and to allow more containers in the future, we create an docker compose file. Our docker compose file looks as follows.
```
version: '3.9'

services:
  minitwit:
    image: lind/minitwit:latest
    build:
      context: "."
      dockerfile: "./Dockerfile"
    ports:
      - "8080:8080"
```

We can now run the application in a daemon using: \
`docker-compose up -d`

We can now see that the application is running.
```
lind@pop-os:~/Desktop/DevOps/DevOps$ docker ps -a
CONTAINER ID   IMAGE           COMMAND       CREATED         STATUS                  PORTS      NAMES
b4fb15d62216   lind/minitwit   "/minitwit"   8 seconds ago   Up 7 seconds            8080/tcp   minitwit
```

## Step 3: Describe Distributed Workflow

We created the file `CONTRIBUTE.md`, which specifies how we will collaborate in our group.

The file looks like this:
```
Repository setup:
It is a public repository with all members having collaborator status.

Branching model:
For this project we have decided on using feature branches.

Distributed development workflow:
The majority of the work will be done as mob programming.

Contribution structure:
As most work is done by mob programming most contributions are made by all repository members.

Reviewing/integrating contributions:
For contributions not done by mob programming, a random member is picked as reviewer and is responsible for merging.
```