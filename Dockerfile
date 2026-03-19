# --- STAGE 1: Build ---
# เปลี่ยนจาก 1.22 เป็น 1.26 ตามที่ Error แจ้ง
FROM golang:1.26-alpine AS builder

# ติดตั้ง git เผื่อไว้สำหรับบาง library
RUN apk add --no-cache git

WORKDIR /

COPY go.mod ./
COPY go.sum ./  
RUN go mod download
#RUN go mod init
#RUN go mod tidy

COPY . .

# Build ไฟล์ binary
RUN GOOS=linux GOARCH=amd64 go build -o /main .


# --- STAGE 2: Run ---
FROM alpine:latest
WORKDIR /
# ก๊อปปี้ไฟล์ main มาจาก stage builder
COPY --from=builder /main .

EXPOSE 3000
RUN chmod +x ./main
CMD ["./main"]