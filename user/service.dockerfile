FROM obraun/vss-jenkins as builder
COPY . /app
WORKDIR /app
RUN go build -o user-service main.go

FROM iron/go
EXPOSE 8091
ENTRYPOINT ["/app/user-service"]
