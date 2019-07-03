# Docker Compose

Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. To learn more about all the features of Compose, see the [list of features](https://docs.docker.com/compose/overview/#features).

Using Compose is basically a three-step process:

1. Define your app’s environment with a `Dockerfile` so it can be reproduced anywhere.
2. Define the services that make up your app in `docker-compose.yml` so they can be run together in an isolated environment.
3. Lastly, run `docker-compose up` and Compose will start and run your entire app.


## Docker Compose File

Create docker compose file.

```yml
version: '3.7'
services:
  ui:
    build:
      context: .
      dockerfile: build/ui/Dockerfile
    image: todo-app-ui:latest
    restart: on-failure
    volumes:
      - ./config/nginx/conf.d:/etc/nginx/conf.d
    networks:
      - default
    ports:
      - 80:80
  app:
    build:
      context: .
      dockerfile: build/app/Dockerfile
    image: todo-app:latest
    restart: on-failure
    command: [
      "--database-host", "database",
      "--database-user", "todoapp",
      "--database-password", "secret",
      "--database-name", "todoapp"
    ]
    networks:
      - default
  database:
    image: mysql:8.0
    volumes:
      - ./data/schema:/docker-entrypoint-initdb.d
      - database-data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: supersecret
      MYSQL_DATABASE: todoapp
      MYSQL_USER: todoapp
      MYSQL_PASSWORD: secret
      TZ: Asia/Jakarta
    networks:
      - default
volumes:
  database_data:
networks:
  default:
```

See [compose file references](https://docs.docker.com/compose/compose-file/) for more details

## Docker Compose CLI

```
Define and run multi-container applications with Docker.

Usage:
  docker-compose [-f <arg>...] [options] [COMMAND] [ARGS...]
  docker-compose -h|--help

Options:
  -f, --file FILE             Specify an alternate compose file
                              (default: docker-compose.yml)
  -p, --project-name NAME     Specify an alternate project name
                              (default: directory name)
  --verbose                   Show more output
  --log-level LEVEL           Set log level (DEBUG, INFO, WARNING, ERROR, CRITICAL)
  --no-ansi                   Do not print ANSI control characters
  -v, --version               Print version and exit
  -H, --host HOST             Daemon socket to connect to

  --tls                       Use TLS; implied by --tlsverify
  --tlscacert CA_PATH         Trust certs signed only by this CA
  --tlscert CLIENT_CERT_PATH  Path to TLS certificate file
  --tlskey TLS_KEY_PATH       Path to TLS key file
  --tlsverify                 Use TLS and verify the remote
  --skip-hostname-check       Don't check the daemon's hostname against the
                              name specified in the client certificate
  --project-directory PATH    Specify an alternate working directory
                              (default: the path of the Compose file)
  --compatibility             If set, Compose will attempt to convert deploy
                              keys in v3 files to their non-Swarm equivalent

Commands:
  build              Build or rebuild services
  bundle             Generate a Docker bundle from the Compose file
  config             Validate and view the Compose file
  create             Create services
  down               Stop and remove containers, networks, images, and volumes
  events             Receive real time events from containers
  exec               Execute a command in a running container
  help               Get help on a command
  images             List images
  kill               Kill containers
  logs               View output from containers
  pause              Pause services
  port               Print the public port for a port binding
  ps                 List containers
  pull               Pull service images
  push               Push service images
  restart            Restart services
  rm                 Remove stopped containers
  run                Run a one-off command
  scale              Set number of containers for a service
  start              Start services
  stop               Stop services
  top                Display the running processes
  unpause            Unpause services
  up                 Create and start containers
  version            Show the Docker-Compose version information
```

You can also see this information by running `docker-compose --help` from the command line.

Create and start docker container

```bash
$ docker-compose -f example/todos-app/todos-app-stack.yml --compatibility up -d
```

Show running docker container

```bash
$ docker-compose -f example/todos-app/todos-app-stack.yml --compatibility ps
```

Show log docker container

```bash
$ docker-compose -f example/todos-app/todos-app-stack.yml --compatibility logs
```

Clean up stopping and removing docker container

```bash
$ docker-compose -f example/todos-app/todos-app-stack.yml --compatibility down
```

## Exercise

- Create docker compose file for `hello-docker` application.