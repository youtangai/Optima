IMAGENAME := optima_worker

main: main.go
	GOOS=linux GOARCH=amd64 go build main.go

clean:
	rm -f main

test:
	go run main.go

build: main
	docker build -t $(IMAGENAME):latest .

run: 
	docker run -d --name $(IMAGENAME) $(IMAGENAME)

down:
	docker stop $(IMAGENAME)
	docker rm $(IMAGENAME)

logs:
	docker logs -f $(IMAGENAME)