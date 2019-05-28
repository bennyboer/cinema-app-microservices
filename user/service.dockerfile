FROM obraun/vss-jenkins as builder
COPY ./user /user-service-app
WORKDIR /user-service-app
RUN go build -o user-service main.go

FROM iron/go
EXPOSE 8091
ENTRYPOINT ["/user-service-app/user-service"]
