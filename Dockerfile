# Step 1: Build the binary
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
# ถ้ามี go.sum ให้เอาคอมเมนต์ออก
COPY go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Step 2: Run the binary
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Azure App Service จะส่ง PORT มาให้เอง
EXPOSE 8080
CMD ["./main"]