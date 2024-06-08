FROM golang:alpine AS backend

# RUN apk add --no-cache build-base

WORKDIR /app

COPY backend ./

RUN go mod download && \
    GOOS=linux GOARCH=amd64 go build -o /out/ai-spellcheck

FROM alpine:latest

WORKDIR /app

COPY --from=backend /out/ai-spellcheck ./

EXPOSE 3000

CMD ["./ai-spellcheck"]