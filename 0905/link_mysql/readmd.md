# Usage

```
docker run -d -e MYSQL_ROOT_PASSWORD=root --name mysql-db mysql
docker run -d -e MYSQL_ROOT_PASSWORD=root --name mysql-client --link mysql-db:db mysql

# enter to db
docker exec -it mysql-db /bin/bash

# in db conainer
mysql -uroot -proot

> create database temp;
> create table temp value (temp varchar(255));

# enter to client
docker exec -it mysql-client /bin/bash

# in client container
mysql -d db -uroot -proot
```
