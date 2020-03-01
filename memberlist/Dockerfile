FROM golang:1.13.1 as builder

ENV GO111MODULE=on
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /src/donkey
COPY ./app.go ./
RUN pwd
RUN ls -al

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o donkey  app.go

FROM alpine:3.11.3 as scratch
WORKDIR /root

COPY --from=builder /src/donkey/donkey .

CMD ["./donkey"]
