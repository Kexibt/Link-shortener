FROM golang:latest as builder 
LABEL maintainer = "Some maintainer kex"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Starting a new stage from scratch 
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 3000 3000

# launch of the project
CMD ["./main"] 
# CMD ["./main -db"]