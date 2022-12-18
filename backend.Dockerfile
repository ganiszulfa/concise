# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

# install dependencies
RUN mkdir -p /app
WORKDIR /app
COPY ./backend/go.mod /app
COPY ./backend/go.sum /app
RUN go mod download

# Copy all local files into the image.
COPY ./backend /app

RUN go build -o /backend

EXPOSE 8080

CMD [ "/backend" ]