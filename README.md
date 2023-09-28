To get started, follow these steps:

1. Please create the .env file with the following command line.
```shell
copy .env.example .env
```

2. Please fill in the .env file with the required configurations below.
```
POSTGRES_USER=
POSTGRES_PASSWORD=
POSTGRES_DB=
```

3. Please start server
```shell
docker compose up
```

4. If you are `dial Tcp 127.0.0.1:5432: If you receive the connect: .... connection refused` error message, type the following command line into the terminal screen and write the `IPAddress` part of the output into the .env file.
```
docker container inspect api-fiber-gorm-database-1
```