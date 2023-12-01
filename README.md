# snippetbox

A basic web application in go 

## Local postgres database setup

### Connect pgadmin client docker container to local postgres server
- Spin up a pgadmin container and port forward to port 5050 of host machine
```sh
docker run -d --name pgadmin --env="PGADMIN_DEFAULT_EMAIL=foo@mail.com" --env="PGADMIN_DEFAULT_PASSWORD=bar" -p 5050:80 dpage/pgadmin4:7
```

- Use browser and connect to pgadmin UI at localhost:5050

- To connect to local postgres server use the special DNS name `host.docker.internal` (which resolves to the internal IP address used by the host) instead of `localhost`

