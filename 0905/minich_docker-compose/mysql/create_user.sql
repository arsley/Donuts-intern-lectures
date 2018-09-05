CREATE USER 'minich_local_user'@'%' IDENTIFIED BY 'minich_local_password';
GRANT ALL ON minich_local.* TO 'minich_local_user'@'%';