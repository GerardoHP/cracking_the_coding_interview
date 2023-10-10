FROM golang:1.20-alpine as build
WORKDIR /app
COPY go.??? /app/
RUN go mod download
COPY . /app/
RUN go mod tidy && \
    go build -o cli ./cmd/cli/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/cli /app/
RUN chmod +x /app/cli
ENV PATH=$PATH:/app
CMD cli