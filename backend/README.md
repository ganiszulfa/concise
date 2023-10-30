#  Concise Backend

## Requirements
1. Go
2. Postgres
2. GraphQL

## Local Deployment

1. Run Postgres
```
$ docker run --name concise-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=concise POSTGRES_USER=postgres -p 5432:5432 -d postgres

```
2. Run Go 
```
go run main.go
```


## Cloud Deployment

1. Concise has been tailored for deployment on [CapRover](https://caprover.com/). CapRover is an amazing piece of software, it's like Heroku of your own! Install this on your own server.
    - Add an app in caprover.
    - Set deployment method to "Official CLI", generate app token.


2. Go to your `Secrets and variables` section in your Github Repo. Add these variables
    - CAPROVER_SERVER
    - APP_NAME
    - APP_TOKEN

3. Push a tag with this format `bv[0-9]+.[0-9]+.[0-9]+`, for example: `bv1.0.2`