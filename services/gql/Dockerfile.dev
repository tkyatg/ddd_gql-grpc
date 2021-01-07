FROM golang:1.13-alpine
WORKDIR /app
RUN apk add --no-cache tzdata git && \
    go get github.com/pilu/fresh

COPY . .

CMD ["fresh"]