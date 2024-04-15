FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cwgo-open-analysis .


FROM alpine

WORKDIR /src
COPY --from=build /app/cwgo-open-analysis /app/cwgo-open-analysis
COPY default.yaml /src/default.yaml

ENTRYPOINT ["/app/cwgo-open-analysis"]