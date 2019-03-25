FROM golang:1.12 as build
WORKDIR /app
COPY . .
RUN go build
CMD ["/app/names"]
