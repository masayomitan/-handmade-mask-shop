FROM golang:1.15.7-alpine

COPY go.mod ./go.mod
COPY go.sum ./go.sum

ENV ROOT=/go/src/app
ENV CGO_ENABLED 0
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
EXPOSE 8080

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

CMD ["go", "run", "main.go", "/go/src/app/binary"]



FROM nginx:latest

COPY ./docker/nginx/default.conf /etc/nginx/nginx.conf