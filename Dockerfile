FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app ./

FROM alpine:3.22
RUN apk add --no-cache ca-certificates \
    && adduser -D -u 10001 appuser
COPY --from=build /app /app
USER appuser
EXPOSE 8080
ENTRYPOINT ["/app"]
