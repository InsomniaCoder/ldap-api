FROM golang:1.16 as builder

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy the go source
COPY main.go main.go
COPY controllers/ controllers/
COPY ldap/ ldap/
COPY mail/ mail/
COPY models/ models/
COPY routes/ routes/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ldap-api main.go

# Use Alpibe as base image since we need slappasswd binary
FROM alpine:3.14.0
WORKDIR /app
COPY --from=builder /workspace/ldap-api .
USER 65532:65532

ENTRYPOINT ["/app/ldap-api"]

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# FROM gcr.io/distroless/static:nonroot
# WORKDIR /app
# COPY --from=builder /workspace/manager .
# USER 65532:65532

# ENTRYPOINT ["/app/manager"]
