# Docker Swarm

## What is a swarm?

The cluster management and orchestration features embedded in the Docker Engine are built using swarmkit. [docker swarm feature](https://docs.docker.com/engine/swarm/#feature-highlights)

### Swarm CLI commands

List available docker swarm command line

```bash
$ docker swarm --help

Usage:	docker swarm COMMAND

Manage Swarm

Commands:
  ca          Display and rotate the root CA
  init        Initialize a swarm
  join        Join a swarm as a node and/or manager
  join-token  Manage join tokens
  leave       Leave the swarm
  unlock      Unlock swarm
  unlock-key  Manage the unlock key
  update      Update the swarm

Run 'docker swarm COMMAND --help' for more information on a command.
```

Initialize a swarm.

```bash
$ docker swarm init
Swarm initialized: current node (dmttjkbgxzl3yb2hsfimq1wr5) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token SWMTKN-1-0yjar5plzt154od7x9zckugo7r7gjxkv8emmrpl9oohn5w044c-f5722yshn0zse4wdymsgxck70 192.168.65.3:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

## Node

A node is an instance of the Docker engine participating in the swarm. You can also think of this as a Docker node. You can run one or more nodes on a single physical computer or cloud server, but production swarm deployments typically include Docker nodes distributed across multiple physical and cloud machines.

### Docker node CLI commands

```bash
$ docker node --help

Usage:	docker node COMMAND

Manage Swarm nodes

Commands:
  demote      Demote one or more nodes from manager in the swarm
  inspect     Display detailed information on one or more nodes
  ls          List nodes in the swarm
  promote     Promote one or more nodes to manager in the swarm
  ps          List tasks running on one or more nodes, defaults to current node
  rm          Remove one or more nodes from the swarm
  update      Update a node

Run 'docker node COMMAND --help' for more information on a command.
```

List nodes in the swarm.

```bash
$ docker node ls
ID                            HOSTNAME                STATUS              AVAILABILITY        MANAGER STATUS      ENGINE VERSION
dmttjkbgxzl3yb2hsfimq1wr5 *   linuxkit-025000000001   Ready               Active              Leader              18.09.2
```

## Service

In a distributed application, different pieces of the app are called “services”. For example, if you imagine a video sharing site, it probably includes a service for storing application data in a database, a service for video transcoding in the background after a user uploads something, a service for the front-end, and so on.

Services are really just “containers in production.” A service only runs one image, but it codifies the way that image runs—what ports it should use, how many replicas of the container should run so the service has the capacity it needs, and so on. Scaling a service changes the number of container instances running that piece of software, assigning more computing resources to the service in the process.

### Deploy hello-docker application with docker service

Create hell_net network with overlay driver

```bash
$ docker network create hello_net --driver overlay
```

Create `web-server` service

```bash
$ docker service create \
  --name web-server \
  --publish 80:80 \
  --replicas 1 \
  --network hello_net \
  --mount type=bind,src=$(pwd)/example/hello-docker/nginx/config/conf.d,dst=/etc/nginx/conf.d \
  --mount type=bind,src=$(pwd)/example/hello-docker/nginx/html,dst=/usr/share/nginx/html \
  nginx:latest
```

Create `hello-go` service

```bash
$ docker service create \
  --name hello-go \
  --replicas 1 \
  --network hello_net \
  hello-go:latest
```

List docker service

```bash
$ docker service ls
```

Logging docker service

```bash
$ docker service logs hello-go
```

Scale `hello-go` docker service

```bash
$ docker service scale hello-go=3
hello-go scaled to 3
overall progress: 3 out of 3 tasks
1/3: running   [==================================================>]
2/3: running   [==================================================>]
3/3: running   [==================================================>]
verify: Service converged
```

## Stack

A stack is a group of interrelated services that share dependencies, and can be orchestrated and scaled together. A single stack is capable of defining and coordinating the functionality of an entire application (though very complex applications may want to use multiple stacks).

Deploy `todo-app` application in docker swarm mode using docker stack

```bash
$ docker-compose -f example/todo-app/todo-app-stack.yml build
$ docker stack deploy -c example/todo-app/todo-app-stack.yml todo-app
```

List docker stack applications

```bash
$ docker stack ls
NAME                SERVICES            ORCHESTRATOR
todo-app            3                   Swarm
```

List docker service in todoapp stack

```bash
$ docker stack services todo-app
ID                  NAME                MODE                REPLICAS            IMAGE                PORTS
8b50ntq9j1os        todo-app_ui         replicated          1/1                 todo-app-ui:latest   *:80->80/tcp
hwdejo8a6f28        todo-app_database   replicated          1/1                 mysql:8.0
kw9inqk9aqoo        todo-app_app        replicated          1/1                 todo-app:latest
```