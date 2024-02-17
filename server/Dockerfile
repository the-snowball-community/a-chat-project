FROM golang:1.21

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
# build
RUN CGO_ENABLED=0 GOOS=linux go build -o /snowball-chat

# port
EXPOSE 4000
CMD ["/snowball-chat"]