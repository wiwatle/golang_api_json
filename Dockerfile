# --- STAGE 1: Build ---
FROM golang:1.26-alpine AS builder

# ติดตั้ง build-base กรณีที่มี library บางตัวต้องใช้ C
RUN apk add --no-cache git build-base

WORKDIR /app

# จัดการเรื่อง Module
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# หัวใจสำคัญ: ปิด CGO และ Build เป็น Static Binary
# เพื่อให้รันบน Alpine ได้โดยไม่ถามหา Library ของระบบ
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# --- STAGE 2: Run ---
FROM alpine:latest
# ติดตั้ง CA-Certificates เผื่อกรณี API ต้องเรียก HTTPS ภายนอก
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ก๊อปปี้ไฟล์จาก builder มาไว้ที่ root ของ run stage
COPY --from=builder /app/main .

# เผื่อว่าในโปรเจกต์มีไฟล์ .env หรือ static files ให้ COPY มาด้วย
# COPY --from=builder /app/.env . 

EXPOSE 8080

# มั่นใจว่าไฟล์รันได้
RUN chmod +x ./main

# สั่งรันจาก path ที่ถูกต้อง
CMD ["./main"]