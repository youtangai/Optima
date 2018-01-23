main: main.go
	GOOS=linux GOARCH=amd64 go build main.go

clean:
	rm -f main

test:
	go run main.go

build: main
	docker build -t print_loop:latest .

run:
	docker run -d --name print_loop print_loop

down:
	docker stop print_loop
	docker rm print_loop

logs:
	docker logs -f print_loop