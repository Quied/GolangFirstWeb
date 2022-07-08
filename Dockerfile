FROM golang:latest

COPY . /home/quied/Desktop/Quied-docker

WORKDIR /home/quied/Desktop/Quied-docker

RUN go build main.go
RUN ./main

CMD ["./main"]
