FROM golang

WORKDIR $GOPATH/src/github.com/tcooper8/thermostat-api/
ADD . .
RUN go get github.com/satori/go.uuid
RUN go get github.com/gorilla/mux
RUN go test
RUN go build -o main .

CMD ["./main"]
