FROM obraun/vss-protoactor-jenkins as builder
COPY . /apps
WORKDIR /apps
RUN sh build.sh

FROM iron/go
EXPOSE 8095
ENTRYPOINT ["/apps/user/user-service"]
