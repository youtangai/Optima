monitor: main.go
	GOOS=linux GOARCH=amd64 go build -o monitor main.go

clean: monitor
	rm -f monitor

send: monitor join_script.sh
	scp monitor join_script.sh leave_script.sh root@192.168.64.14:/opt

sendall: monitor join_script.sh
	scp monitor join_script.sh leave_script.sh root@192.168.64.14:/opt
	scp monitor join_script.sh leave_script.sh root@192.168.64.16:/opt
	scp monitor join_script.sh leave_script.sh root@192.168.64.17:/opt

sendleav:
	scp leave_script.sh root@192.168.64.14:/opt
	scp leave_script.sh root@192.168.64.16:/opt
	scp leave_script.sh root@192.168.64.17:/opt