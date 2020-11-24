FROM golang:1.15.5-alpine3.12

COPY ./internal /go/src/github.com/florianwoelki/insta-clone/account-api

WORKDIR /go/src/github.com/florianwoelki/insta-clone/account-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

CMD [ "account-api" ]
