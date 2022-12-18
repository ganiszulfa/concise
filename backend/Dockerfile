# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

# install dependencies
RUN mkdir -p /app
WORKDIR /app
COPY ./go.mod /app
COPY ./go.sum /app
RUN go mod download

# Copy all local files into the image.
COPY ./ /app

RUN go build -o /backend

EXPOSE 8080

CMD [ "/backend" ]