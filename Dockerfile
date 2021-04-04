FROM golang

EXPOSE 8080:8080

COPY . /go/src/gin_msg

WORKDIR /go/src/gin_msg


RUN go build
RUN go install

CMD ["gin_msg"]
