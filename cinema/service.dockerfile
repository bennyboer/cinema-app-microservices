FROM obraun/vss-jenkins as builder
COPY ./user /cinema-service-app
WORKDIR /cinema-service-app
RUN go build -o cinema-service main.go

FROM iron/go
EXPOSE 8091
ENTRYPOINT ["/cinema-service-app/cinema-service"]
