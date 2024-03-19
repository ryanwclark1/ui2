FROM golang:1.22-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy your code into the container.
COPY . .

# Set necessary environment variables and build your project.
ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o default

FROM scratch

# Copy project's binary and templates from /build to the scratch container.
COPY --from=builder /build/default /



# Set entry point.
ENTRYPOINT ["/default"]
