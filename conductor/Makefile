conductor: main.go
	GOOS=linux GOARCH=amd64 go build -o conductor main.go

clean:conductor
	rm -f conductor

send:conductor
	scp conductor root@192.168.64.12:/opt