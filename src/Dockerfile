FROM golang

EXPOSE 8080:8080

COPY . /go/src/asynji

WORKDIR /go/src/asynji


RUN go build
RUN go install

CMD ["asynji"]
