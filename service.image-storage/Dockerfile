FROM golang:1.15.5-alpine3.12 as builder

ADD . /go/src/github.com/florianwoelki/insta-clone/

WORKDIR /go/src/github.com/florianwoelki/insta-clone/service.image-storage

COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /image-storage . -ldflags "-X main.env=docker"

FROM alpine

RUN apk update \
  && apk add --no-cache

COPY --from=builder /account-api ./
RUN chmod +x ./image-storage
CMD [ "./image-storage" ]
