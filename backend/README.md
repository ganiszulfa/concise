# BACKEND

## Local Development

### DB

```
docker run --name pg -e POSTGRES_PASSWORD=postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_DB=concise -d postgres
```