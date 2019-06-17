FROM obraun/vss-jenkins as builder
COPY ./cinema /cinema-service-app
WORKDIR /cinema-service-app
RUN sh git config --global user.name "bennyboer-machine-user"
RUN sh git config --global credential.helper store
RUN sh echo https://51faa31d4b9f08c8e56d4fb23fc082a85e617df8:x-oauth-basic@github.com >> ~/.git-credentials
RUN go build -o cinema-service main.go

FROM iron/go
EXPOSE 8091
ENTRYPOINT ["/cinema-service-app/cinema-service"]
