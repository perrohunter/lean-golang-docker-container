Lean Golang Docker Container
=====

To build this container:

```
$ make
```

This container needs another container from `nats` to be running

```
docker run -it -d --name nats -p 8222:8222 -p 4222:4222 -p 6222:6222 nats:0.9.6
```

Then run this container

```
docker run -p 8080:8080 --link nats lean-golang-docker-container
```