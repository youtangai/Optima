monitor: main.go
	GOOS=linux GOARCH=amd64 go build -o monitor main.go

clean: monitor
	rm -f monitor

send: monitor join_script.sh
	scp monitor root@192.168.64.14:/opt
	scp join_script.sh root@192.168.64.14:/opt