FROM golang:latest
COPY . /opt/app
WORKDIR /opt/app/cmd
RUN go build -o main .
CMD ["./main"]
