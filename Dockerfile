FROM golang:1.13

WORKDIR /app
COPY . .
RUN go get -d github.com/gorilla/mux

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:latest
WORKDIR /app
COPY --from=0 /app/main .

COPY ./public/  /app/public/
COPY ./pages/   /app/pages/
COPY ./themes/  /app/themes/

CMD ["/app/main"]

EXPOSE 8080