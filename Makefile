default: builddocker

setup:
	go get github.com/nats-io/go-nats

buildgo:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o main ./go/src/lean-golang-docker-container

builddocker:
	docker build -t build-image -f ./Dockerfile.Go .
	docker run -t build-image /bin/true
	docker cp `docker ps -q -n=1`:/main .
	chmod 755 ./main
	docker build --rm=true --tag=lean-golang-docker-container -f Dockerfile.static .

run: builddocker
	docker run \
		-p 8080:8080 lean-golang-docker-container