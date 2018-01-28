dbmigration: main.go
	GOOS=linux GOARCH=amd64 go build -o dbmigration main.go

send:dbmigration
	scp dbmigration root@192.168.64.12:/opt

honban:dbmigration
	scp dbmigration root@10.6.18.162:/opt
