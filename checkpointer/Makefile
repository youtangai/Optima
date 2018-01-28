checkpointer: main.go
	GOOS=linux GOARCH=amd64 go build -o checkpointer main.go
clean: checkpointer
	rm -rf checkpointer

send: checkpointer
	scp checkpointer root@192.168.64.14:/opt

sendall: checkpointer
	scp checkpointer root@192.168.64.14:/opt
	scp checkpointer root@192.168.64.16:/opt
	scp checkpointer root@192.168.64.17:/opt

honban: checkpointer
	scp checkpointer root@10.6.18.218:/opt
	scp checkpointer root@10.6.18.166:/opt
