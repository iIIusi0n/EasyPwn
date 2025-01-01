FROM debian:bookworm AS builder

RUN apt-get update && apt-get install -y \
    curl \
    git \
    unzip \
    xz-utils \
    && rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/flutter/flutter.git /flutter
ENV PATH="/flutter/bin:${PATH}"

RUN flutter doctor
RUN flutter config --enable-web

WORKDIR /app
COPY web .

RUN flutter pub get
RUN flutter build web --release

FROM nginx:1.27.3-alpine3.20

COPY --from=builder /app/build/web /usr/share/nginx/html

COPY assets/nginx.conf /etc/nginx/nginx.conf

CMD ["nginx", "-g", "daemon off;"] 