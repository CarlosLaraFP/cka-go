# Uses a multi-stage build – Keeps the final image small and clean.
# Removes Go compiler from the final image – Reduces attack surface.
# Uses Distroless (gcr.io/distroless/base-debian12) – More secure, removes shell access.
# Efficient caching – First copies go.mod & go.sum, then downloads dependencies separately.

# We use Distroless because security is a top priority:

# No unnecessary binaries reduce the attack surface.
# We don’t need a shell or package manager – the Go app runs standalone.
# We want a minimal image – No need for OS utilities, just the app.
# We are deploying to Kubernetes – Distroless aligns well with lightweight, secure containers.
# No package manager (smaller attack surface).
# Minimalist runtime (low overhead).
# Used in production by Google Kubernetes Engine (GKE).

# First stage: Build the binary
FROM golang:1.24 AS builder
WORKDIR /app

# Cache Go modules separately
COPY go.mod go.sum ./
RUN go mod download

# Copy the remaining code and build the binary
COPY . .
RUN go build -o go-app .

# Second stage: Use a minimal base image
#FROM gcr.io/distroless/base-debian12
#COPY --from=builder /app/go-app /go-app

#CMD ["/go-app"]
CMD ["/app/go-app"]