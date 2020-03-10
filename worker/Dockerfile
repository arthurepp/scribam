FROM golang:latest AS build-env
RUN go get github.com/streadway/amqp
RUN go get -u github.com/arthurepp/goutil
RUN go get gopkg.in/mgo.v2
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux go build -o main . 

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /app/main .
CMD ["./main"]