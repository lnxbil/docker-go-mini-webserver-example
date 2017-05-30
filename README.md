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
can save the image and load it on the destination docker machine. Be aware
that this is not an `export` and `import` operation, but an `save` and `load`
one:

    make save

to save it copy the file to the remote location's `/tmp`folder and load it
there:

    docker load -i /tmp/docker-go-mini-webserver-example.docker-save

Afterwards, you'll have an image, not an container that can now be used
to create a persistent or throw-away container.

    docker run \
      --interactive --tty --rm --publish 8081:8081 \
      docker-go-mini-webserver-example-app:latest
