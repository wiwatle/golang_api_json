# --- STAGE 1: Build ---
FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./  
RUN go mod download
COPY . .

# *** จุดสำคัญ 1: สั่ง build และตั้งชื่อ output ว่า 'main' ***
# ตรวจสอบว่ามีจุด (.) ต่อท้ายด้วย เพื่อบอกให้ build จากโฟลเดอร์ปัจจุบัน
RUN GOOS=linux GOARCH=amd64 go build -o /app/main .

# --- STAGE 2: Run ---
FROM alpine:latest
WORKDIR /root/

# *** จุดสำคัญ 2: ตรวจสอบ Path ให้ตรงกับ Stage แรก ***
# เราก๊อปปี้จาก builder มาที่พาธปัจจุบัน (.)
COPY --from=builder /app/main .

EXPOSE 8080

# *** จุดสำคัญ 3: สั่งรันไฟล์ที่เราก๊อปมา ***
CMD ["./main"]