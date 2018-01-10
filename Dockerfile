FROM golang:1.9-alpine

ENV MAILGUN_SANDBOX_DOMAIN sandbox3fd66e607c004414a32485674ce9674f.mailgun.org
ENV MAILGUN_API_KEY key-81f9124c51b2a16192b5d546df0cae3b
ENV MAILGUN_PUB_API_KEY pubkey-27244492e5d0f52210d61f36ec18620c
ENV PUBG_TRN_ENDPOINT https://api.pubgtracker.com/v2/profile/pc/adamms
ENV TRN_API_KEY 17d76a22-907b-4ab0-84fe-dfbc34cebd0c    

ADD . $GOPATH/src/github.com/adamfdl/pubg_trn_status
WORKDIR $GOPATH/src/github.com/adamfdl/pubg_trn_status
RUN go build

CMD [ "./pubg_trn_status" ]