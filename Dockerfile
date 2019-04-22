FROM golang:1-alpine
COPY . /go/src/srvbin
WORKDIR /go/src/srvbin
RUN CGO_ENABLED=0 go build main.go

FROM scratch
COPY --from=0 /go/src/srvbin/main /usr/bin/srvbin
ENTRYPOINT ["srvbin"]
