FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o ./myapp dummyservice/dummyservice.go 

FROM alpine:latest
COPY --from=builder /app/myapp /bin/myapp
RUN chmod +x /bin/myapp
EXPOSE 8080
CMD ["/bin/myapp"]