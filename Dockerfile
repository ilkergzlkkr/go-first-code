FROM go:alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 3000

# using ./main.go as the entrypoint

RUN go build -o main .

CMD ["./main"]