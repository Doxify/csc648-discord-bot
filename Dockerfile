FROM golang:alpine
WORKDIR /home/app
COPY . /home/app
RUN go build
CMD ["./csc648-discord-bot"]