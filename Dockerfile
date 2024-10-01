###################
# 1. Building stage
###################
FROM golang:1.23.1 AS build-stage

# Set destination for COPY
WORKDIR /waku-cli

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY *.go ./
COPY cmd/ pkg/ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o waku



###################
# 2. Run tests
###################
FROM build-stage AS run-test-stage
RUN go test -v ./...



###################
# 3. Deploying
###################
FROM alpine:3.20.3 AS deploy-stage
WORKDIR /app

# Install git
RUN apk add --update --no-cache tini git && rm -rf /var/cache/apk/*

RUN adduser -D waku
USER waku

# Copy bins from build stage
COPY --from=build-stage /waku-cli/waku /usr/bin/waku

# Run
ENTRYPOINT ["/sbin/tini", "--", "waku"]
