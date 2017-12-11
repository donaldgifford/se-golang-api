# se-golang-api

A simple golang api that exposes a status endpoint. This API is intended to be connected to a postgres db in which it performs a status check on if the API can connect or not. 

Libraries used:

 - [hootsuite/healthchecks](https://github.com/hootsuite/healthchecks)
 - [lib/pq](https://github.com/lib/pq)

#### Running se-golang-api

Docker repo - [se-golang-api](https://hub.docker.com/r/donaldgifford/se-golang-api/)

ENV VARS to specify DB connections:
```
	"DB_NAME" - "The database used IE - api"
	"DB_USER" - "The postgres user to connect with"
    "DB_PASS" - "The postgres password to connect with"
	"DB_HOST" - "connection string to the db"
	"DB_PORT" - "The postgres db port - typically 5432"
```

Example docker run:
```
docker run -e DB_USER="postgres" \
    -e DB_PASS="example" \
    -e DB_HOST="localhost" \
    -e DB_PORT="5432" \
    -e DB_NAME="api" \
    -p 8000:8000 \
    donaldgifford/se-golang-api
```

Testing locally:
```
docker-compose up
```

This uses the adminer image which is a simple php DB management tool. To run using docker-compose you have to create the database the api is set to connect with. More info [here](https://github.com/vrana/adminer)