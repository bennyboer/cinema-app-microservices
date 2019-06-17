FROM obraun/vss-protoactor-jenkins as builder
COPY . /apps
WORKDIR /apps
RUN sh git config --global user.name "bennyboer-machine-user"
RUN sh git config --global credential.helper store
RUN sh echo https://51faa31d4b9f08c8e56d4fb23fc082a85e617df8:x-oauth-basic@github.com >> ~/.git-credentials
RUN sh build.sh

FROM iron/go
EXPOSE 8095
ENTRYPOINT ["/apps/user/user-service"]
