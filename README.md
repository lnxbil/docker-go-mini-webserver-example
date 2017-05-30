# Minimal Working Example

This minimal working example (MWE) shows how you can create a simple docker
container with a single go binary that provides a webserver running on port
`8081` and is exposed through docker to the outside world, such that you can
access it directly on

    http://localhost:8081/

on your machine that runs the docker container.


# Requirements

Besides the `make` and `docker` command, internet access is also required.

# Manual Distribution

If you do not want to use a registry to distribute the docker image, you
can create a container, export and import it on the destination docker
machine:

    make export

to export it, copy the file to the remote location's `/tmp`folder and import

    docker import /tmp/docker-go-mini-webserver-example.docker-export
