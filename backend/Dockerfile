FROM golang:1.22.5-alpine AS build

ENV GIN_MODE=release
# ENV DB_URL=some-postgres

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod .
COPY go.sum .

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o main .

CMD ["/app/main"]