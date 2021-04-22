FROM golang:1.16-buster as builder

# ygulama dizinini oluşturun ve değiştirin.
WORKDIR /app

# Uygulama bağımlılıklarını yükleyin..
# go.mod ve varsa go.sum kopyalanması gerekiyor.
COPY go.* ./
RUN go mod download

# local kodu konteynar imagesine kopyalayın.
COPY . ./

# binary oluşturun
RUN go build -v -o server

# Yalın bir production konteynar için resmi Debian slim image kullanın.
# https://hub.docker.com/_/debian
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# builder stagedeki binary dosyasını production imageye kopyalayın
COPY --from=builder /app/server /app/server

# Web servisi konteynar başlangıcında çalıştırın.
CMD ["/app/server"]