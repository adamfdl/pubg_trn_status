FROM golang:1.9-alpine

ADD . $GOPATH/src/github.com/adamfdl/pubg_trn_status
WORKDIR $GOPATH/src/github.com/adamfdl/pubg_trn_status
RUN go build

CMD [ "./pubg_trn_status" ]