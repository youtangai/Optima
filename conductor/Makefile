IMAGE_NAME=optima_conductor

main: main.go
	GOOS=linux GOARCH=amd64 go build main.go

clean:
	rm -f main

test:
	go run main.go

build: main
	docker build -t $(IMAGE_NAME):latest .

run:
	docker run -d --name $(IMAGE_NAME) $(IMAGE_NAME)

down:
	docker stop $(IMAGE_NAME)
	docker rm $(IMAGE_NAME)

logs:
	docker logs -f $(IMAGE_NAME)