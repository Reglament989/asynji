FROM golang:1.14

ENV PORT 3000
EXPOSE 3000:8081

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
