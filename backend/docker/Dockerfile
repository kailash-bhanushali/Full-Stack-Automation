FROM golang:1.17.2 as build-image

ENV GOPATH=/go \
    HOME=/go \
    REPO=demo-backend



COPY . $HOME/src/$REPO
RUN cd /$HOME/src/$REPO && make build
WORKDIR /$HOME/bin

#COPY ./entrypoint.sh /usr/bin/entrypoint.sh
#RUN chmod +x /usr/bin/entrypoint.sh
#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN cd /$HOME/src/$REPO
#WORKDIR /$HOME/bin
#RUN go build
EXPOSE 8080
ENTRYPOINT ["bash", "entrypoint.sh"]