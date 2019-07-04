# Docker Machine

## What is Docker Machine?

Docker Machine is a tool that lets you install Docker Engine on virtual hosts, and manage the hosts with docker-machine commands. You can use Machine to create Docker hosts on your local Mac or Windows box, on your company network, in your data center, or on cloud providers like Azure, AWS, or Digital Ocean.

Using docker-machine commands, you can start, inspect, stop, and restart a managed host, upgrade the Docker client and daemon, and configure a Docker client to talk to your host.

## What’s the difference between Docker Engine and Docker Machine?

When people say “Docker” they typically mean Docker Engine, the client-server application made up of the Docker daemon, a REST API that specifies interfaces for interacting with the daemon, and a command line interface (CLI) client that talks to the daemon (through the REST API wrapper).

**Docker Machine** is a tool for provisioning and managing your Dockerized hosts (hosts with Docker Engine on them). Typically, you install Docker Machine on your local system. Docker Machine has its own command line client docker-machine and the Docker Engine client, docker. You can use Machine to install Docker Engine on one or more virtual systems. These virtual systems can be local (as when you use Machine to install and run Docker Engine in VirtualBox on Mac or Windows) or remote (as when you use Machine to provision Dockerized hosts on cloud providers). The Dockerized hosts themselves can be thought of, and are sometimes referred to as, managed “machines”.

## Get Started with Docker Machine

List available docker machine

```bash
$ docker-machine ls
NAME   ACTIVE   DRIVER   STATE   URL   SWARM   DOCKER   ERRORS
```

Create docker machine called machine-0 on localhost using virtualbox

```bash
$ docker-machine create --driver virtualbox machine-0
Running pre-create checks...
Creating machine...
(machine-0) Copying .docker/machine/cache/boot2docker.iso to .docker/machine/machines/machine-0/boot2docker.iso...
(machine-0) Creating VirtualBox VM...
(machine-0) Creating SSH key...
(machine-0) Starting the VM...
(machine-0) Check network to re-create if needed...
(machine-0) Found a new host-only adapter: "vboxnet0"
(machine-0) Waiting for an IP...
Waiting for machine to be running, this may take a few minutes...
Detecting operating system of created instance...
Waiting for SSH to be available...
Detecting the provisioner...
Provisioning with boot2docker...
Copying certs to the local machine directory...
Copying certs to the remote machine...
Setting Docker configuration on the remote daemon...
Checking connection to Docker...
Docker is up and running!
To see how to connect your Docker Client to the Docker Engine running on this virtual machine, run: docker-machine env machine-0
```

Verify new docker machine

```bash
$ docker-machine ls
NAME        ACTIVE   DRIVER         STATE     URL                          SWARM   DOCKER     ERRORS
machine-0   -        virtualbox     Running   tcp://192.168.99.100:2376            v18.09.6
```

Configure shell to use docker engine from docker machine `machine-0`

```bash
$ docker-machine env machine-0
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.100:2376"
export DOCKER_CERT_PATH="/Users/prksu/.docker/machine/machines/machine-0"
export DOCKER_MACHINE_NAME="machine-0"
# Run this command to configure your shell:
# eval $(docker-machine env machine-0)
```

*note: this is only affected the current shell*


## Docker Machine CLI

List available docker machine command line

```bash
Usage: docker-machine [OPTIONS] COMMAND [arg...]

Create and manage machines running Docker.

Version: 0.16.1, build cce350d7

Author:
  Docker Machine Contributors - <https://github.com/docker/machine>

Options:
  --debug, -D                   Enable debug mode
  --storage-path, -s            Configures storage path [$MACHINE_STORAGE_PATH]
  --tls-ca-cert                 CA to verify remotes against [$MACHINE_TLS_CA_CERT]
  --tls-ca-key                  Private key to generate certificates [$MACHINE_TLS_CA_KEY]
  --tls-client-cert             Client cert to use for TLS [$MACHINE_TLS_CLIENT_CERT]
  --tls-client-key              Private key used in client TLS auth [$MACHINE_TLS_CLIENT_KEY]
  --github-api-token            Token to use for requests to the Github API [$MACHINE_GITHUB_API_TOKEN]
  --native-ssh                  Use the native (Go-based) SSH implementation. [$MACHINE_NATIVE_SSH]
  --bugsnag-api-token           BugSnag API token for crash reporting [$MACHINE_BUGSNAG_API_TOKEN]
  --help, -h                    show help
  --version, -v                 print the version

Commands:
  active            Print which machine is active
  config            Print the connection config for machine
  create            Create a machine
  env               Display the commands to set up the environment for the Docker client
  inspect           Inspect information about a machine
  ip                Get the IP address of a machine
  kill              Kill a machine
  ls                List machines
  provision         Re-provision existing machines
  regenerate-certs  Regenerate TLS Certificates for a machine
  restart           Restart a machine
  rm                Remove a machine
  ssh               Log into or run a command on a machine with SSH.
  scp               Copy files between machines
  mount             Mount or unmount a directory from a machine with SSHFS.
  start             Start a machine
  status            Get the status of a machine
  stop              Stop a machine
  upgrade           Upgrade a machine to the latest version of Docker
  url               Get the URL of a machine
  version           Show the Docker Machine version or a machine docker version
  help              Shows a list of commands or help for one command

Run 'docker-machine COMMAND --help' for more information on a command.
```

## Exercise

- Create another docker-machine using virtualbox driver with name `machine-1`
- Verify `machine-1` was created using `docker-machine ls` command.