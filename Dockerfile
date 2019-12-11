FROM golang:1.12 as build
WORKDIR /app
COPY . .
RUN go build -o names
RUN ls -la /app
CMD ["/app/names"]
