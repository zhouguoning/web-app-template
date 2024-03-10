FROM golang:1.13

WORKDIR /
COPY . .
RUN go get -d github.com/gorilla/mux

CMD ["go","run","main.go"]

EXPOSE 8080