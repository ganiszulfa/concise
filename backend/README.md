#  Concise Backend 

## Tech Specs
1. Go
2. Gorm
2. Postgres
2. GraphQL

## Local Deployment

1. (Optional if you don't have Postgres) Run Postgres
```
$ docker run --name concise-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=concise POSTGRES_USER=postgres -p 5432:5432 -d postgres

```
2. Copy `.env.example` to `.env` and adjust accordingly.
2. Run the backend service 
```
go run main.go
```
4. Go to http://localhost:8080/health to check it.

#### Optional:

- Use Air for Live Reload during development (https://github.com/cosmtrek/air)


## Cloud Deployment

1. Concise Backend has been tailored for deployment on [CapRover](https://caprover.com/). CapRover is an amazing piece of software, it's like Heroku of your own! Install this on your own server.
    - (Optional) Create a postgres DB in CapRover from the list.
    - Add a custom app in CapRover (the name will be BE_APP_NAME).
    - Set deployment method to "Official CLI", generate app token (it will be BE_APP_TOKEN).
    - Go to App Configs section, set these values in `Environmental Variables`
```
CON_DEBUG_MODE=false
CON_DB_USER=[ADD YOUR OWN VALUE HERE]
CON_DB_PW=[ADD YOUR OWN VALUE HERE]
CON_DB_PORT=[ADD YOUR OWN VALUE HERE]
CON_DB_HOST=[ADD YOUR OWN VALUE HERE]
CON_DB_NAME=[ADD YOUR OWN VALUE HERE]
```


2. Go to your `Secrets and variables` section in your Github Repo. Add these variables
    - BE_CAPROVER_SERVER
    - BE_APP_NAME
    - BE_APP_TOKEN

3. Push a tag with this format `bv[0-9]+.[0-9]+.[0-9]+`, for example: `bv1.0.2`