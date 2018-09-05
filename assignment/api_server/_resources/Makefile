NAME    := minich-api-server

fmt:
	go fmt

local: fmt
	go build  -o=bin/$(NAME)-local

prepare-local:
	cat create_user.sql | docker exec -i mysql mysql -h 127.0.0.1 -uroot -pmysql

reset:
	echo 'DROP DATABASE IF EXISTS minich_local; CREATE DATABASE minich_local;' | docker exec -i mysql mysql -h 127.0.0.1 -uroot -pmysql
	cat schema.sql | docker exec -i mysql mysql -h 127.0.0.1 -uroot -pmysql minich_local

prepare-exercise:
	cat prepare.sql | docker exec -i mysql mysql -h 127.0.0.1 -uroot -pmysql minich_local