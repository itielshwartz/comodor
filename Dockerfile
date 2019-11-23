FROM golang:latest AS builder
ADD . /app/
WORKDIR /app/agent
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./app/agent/main
RUN chmod +x ./app/agent/main
ENTRYPOINT ["./app/agent/main"]
EXPOSE 3030
