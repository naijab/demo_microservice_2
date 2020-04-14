FROM golang:alpine AS buid-env
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/post-service

FROM scratch
COPY --from=buid-env /go/bin/post-service /go/bin/post-service
ENTRYPOINT ["/go/bin/post-service"]

EXPOSE 1332