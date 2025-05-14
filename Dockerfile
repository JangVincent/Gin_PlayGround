# Go 공식 이미지 사용
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 필요한 패키지 설치 (git 등)
RUN apk add --no-cache git

# go mod 파일 복사 및 의존성 설치
COPY go.mod go.sum ./
RUN go mod download

# 소스 코드 복사
COPY . .

# 빌드
RUN go build -o server main.go

# --- 런타임 이미지 ---
FROM alpine:latest
WORKDIR /app

# # CA 인증서 등 설치 (Postgres TLS 등 필요시)
# RUN apk add --no-cache ca-certificates

# 빌드된 바이너리 복사
COPY --from=builder /app/server .
COPY .env .

# 8080 포트 오픈
EXPOSE 3000

# 서버 실행
CMD ["./server"] 
