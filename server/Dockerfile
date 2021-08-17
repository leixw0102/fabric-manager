FROM golang:1.16
WORKDIR /fabric-manager/server
COPY . .
RUN go mod tidy
RUN go build -o server
CMD [ "./server" ]