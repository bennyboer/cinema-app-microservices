FROM obraun/vss-jenkins as builder
COPY ./cinema /cinema-service-app
WORKDIR /cinema-service-app
RUN go build -o cinema-service main.go

FROM iron/go
EXPOSE 8096
ENTRYPOINT ["/cinema-service-app/cinema-service"]
