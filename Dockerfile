FROM golang:1.23.1 AS build

WORKDIR /eterion_around_us

COPY go.mod go.sum /eterion_around_us/
RUN go mod download

COPY . .

RUN go build -o ./eterion_around_us_build

FROM alpine:3.12 AS run

RUN apk --no-cache add curl

COPY --from=build /eterion_around_us/eterion_around_us_build /eterion_around_us/eterion_around_us_build

WORKDIR /eterion_around_us

EXPOSE 3000

CMD ["./eterion_around_us_build"]