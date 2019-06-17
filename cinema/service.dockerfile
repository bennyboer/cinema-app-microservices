FROM obraun/vss-protoactor-jenkins as builder
RUN mkdir -p /go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/
COPY . /go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/
WORKDIR /go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema
RUN sh build.sh

FROM iron/go
ENTRYPOINT ["/go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/cinema/cinema-service"]