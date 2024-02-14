FROM golang:1.21

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /snowball-chat

EXPOSE 4000
CMD ["/snowball-chat"]