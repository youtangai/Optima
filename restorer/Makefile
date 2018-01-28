restorer: main.go
	GOOS=linux GOARCH=amd64 go build -o restorer main.go
clean: restorer
	rm -rf restorer

send: restorer
	scp restorer root@192.168.64.14:/opt

sendall: restorer
	scp restorer root@192.168.64.14:/opt
	scp restorer root@192.168.64.16:/opt
	scp restorer root@192.168.64.17:/opt

honban: restorer
	scp restorer root@10.6.18.218:/opt
	scp restorer root@10.6.18.166:/opt
