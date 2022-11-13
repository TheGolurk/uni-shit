FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN ls

COPY . .
RUN go build -o app cmd/main.go

EXPOSE 8070

CMD ["./app"]