# Golang Authentication API with Fiber MongoDB and JWT

## Run Database on Docker

```bash
  docker run -it --rm --name mongodb_container -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin -v mongodata:/data/db -d -p 27017:27017 mongo

    docker exec -it mongodb_container /bin/bash

    mongo -u admin -p admin --authenticationDatabase admin

    use mydb

    db.createUser({user: 'user', pwd: 'password', roles:[{'role': 'readWrite', 'db': 'mydb'}]});

    # testing authentication with new user
    mongo -u user -p password --authenticationDatabase mydb

    use mydb

    show collections
```

## Create Your Env File

```bash
DATABASE_USER=user
DATABASE_PASS=password
DATABASE_HOST=127.0.0.1
DATABASE_PORT=27017
DATABASE_NAME=mydb
JWT_SECRET_KEY=secret
```

## Run API

```bash
go run main.go
```

### References:

https://github.com/gofiber/fiber

https://sodocumentation.net/go/topic/10161/jwt-authorization-in-go