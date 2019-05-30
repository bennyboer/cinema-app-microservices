FROM obraun/vss-protoactor-jenkins as builder
COPY ./movie /movie-service-app
WORKDIR /movie-service-app
RUN go build -o movie-service main.go

FROM iron/go
EXPOSE 8093
ENTRYPOINT ["/movie-service-app/movie-service"]
