build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o netatmo2influx .
dockerbuild:
	docker build -t netatmo2influx -f Dockerfile .
dockerrun:
	docker run -it netatmo2influx
distribute:
	docker tag netatmo2influx 192.168.1.231:5000/netatmo2influx
	docker push 192.168.1.231:5000/netatmo2influx
run:
	go run main.go
dockerclean:
	docker rm $(shell docker ps --all -q -f status=exited)
