# ---------- Build Stage ----------
FROM golang as builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY go.* ./
RUN go mod download

# Copy all source files
COPY . ./

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server

# ---------- Final Stage ----------
FROM scratch

# Copy binary from builder
COPY --from=builder /app/server /server

# Copy .env if your app strictly requires it at runtime
COPY --from=builder /app/.env /.env

# Set binary entrypoint
ENTRYPOINT ["/server"]
