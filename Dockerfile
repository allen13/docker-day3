FROM golang:1.15.1-alpine AS builder

COPY hello-http.go .

ENV GO111MODULE="on" \
    CGO_ENABLED=0 \
    GOOS=linux

RUN go build -ldflags="-s -w" -o hello-http ./hello-http.go

# step 2 - run image
FROM scratch AS runner
# Copy our static executable.
COPY --from=builder /go/hello-http /usr/bin/hello-http
