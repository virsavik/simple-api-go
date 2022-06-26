# syntax=docker/dockerfile:1

FROM golang:1.18.3-alpine3.16

WORKDIR /opt

COPY . .

RUN go build -o /app cmd/serverd/main.go

FROM golang:1.18.3-alpine3.16

WORKDIR /opt

COPY --from=0 /app .

EXPOSE 8000

CMD [ "/opt/app" ]