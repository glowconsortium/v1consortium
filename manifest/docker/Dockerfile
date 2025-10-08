# Build stage
FROM golang:1.25-alpine AS builder

# Install required packages for building
RUN apk add --no-cache git wget

# Set working directory
WORKDIR /src

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Install GoFrame CLI
RUN wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_linux_amd64 && \
    chmod +x gf && \
    ./gf install -y && \
    rm ./gf

# Copy source code
COPY . .

# Build the application using GoFrame CLI
RUN gf build -a amd64 -s linux -p temp -ew

# Runtime stage
FROM alpine:3.18

###############################################################################
#                                INSTALLATION
###############################################################################

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

ENV WORKDIR /app

# Copy built binary and resources
COPY --from=builder /src/temp/linux_amd64/main $WORKDIR/main
COPY --from=builder /src/manifest/config/config.yaml $WORKDIR/config.yaml

COPY resource $WORKDIR/

RUN chmod +x $WORKDIR/main

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
