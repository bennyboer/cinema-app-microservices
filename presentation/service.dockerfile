FROM obraun/vss-protoactor-jenkins as builder
RUN mkdir -p /go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/
COPY . /go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/
WORKDIR /go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation
RUN sh build.sh docker

FROM iron/go
ENTRYPOINT ["/go/src/github.com/ob-vss-ss19/blatt-4-sudo_blatt4/presentation/presentation-service"]