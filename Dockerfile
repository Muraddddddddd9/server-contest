FROM golang:latest

WORKDIR /

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]