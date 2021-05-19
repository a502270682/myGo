FROM  golang:latest

#ENV GO111MODULE on
#WORKDIR /test/cache
#ADD go.mod .
#ADD go.sum .
#RUN go mod download
#WORKDIR /test/release
#ADD build .
#RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o myGo main.go

WORKDIR /root
COPY /myGo /
COPY /config/mygo.toml /etc/config/mygo.toml

EXPOSE 8080
CMD ["/myGo", "-c", "/etc/config/mygo.toml"]
