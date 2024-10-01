FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go test ./...
RUN go build -o shipment-calculator main.go

ENV PORT=80
ENV PACK_SIZES_LOCATION=./priv/sizes.json

CMD ["./shipment-calculator", "-p", "80", "-l", "./priv/sizes.json", "start"]