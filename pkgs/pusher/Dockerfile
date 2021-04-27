FROM golang

# EXPOSE 8080:8080

COPY . /go/src/pusher

WORKDIR /go/src/pusher


RUN go build
RUN go install

CMD ["pusher"]
