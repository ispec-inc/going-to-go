# {{cookiecutter.repository_name}}

## Usage
1. run server

Docker
```
$ make migrate
$ make run
```

Local
```
$ mysql -u root
mysql> create database {{cookiecutter.db_name}}

$ dbmate up
$ go run github.com/{{cookiecutter.organization_name}}/{{cookiecutter.repository_name}}
```

2. healthcheck

```
$ curl localhost:9000/health
{"message":"success"}
```
