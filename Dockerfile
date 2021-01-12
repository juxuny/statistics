FROM golang:1.13.4 AS builder
COPY . /src
RUN cd /src && go env -w GOPROXY=https://goproxy.cn && go mod download
RUN mkdir /app
RUN cd /src/cmd/collect &&  CGO_ENABLED=0 go build -o /app/collect
RUN cd /src/cmd/dl-history && CGO_ENABLED=0 go build -o /app/dl-history
RUN cd /src/cmd/dl-sohu && CGO_ENABLED=0 go build -o /app/dl-sohu
RUN cd /src/cmd/statistics-cli && CGO_ENABLED=0 go build -o /app/statistics-cli
RUN cd /src/cmd/stock-code && CGO_ENABLED=0 go build -o /app/stock-code

# final stage
FROM ineva/alpine:3.9
WORKDIR /app
COPY --from=builder /app/collect /app/collect
COPY --from=builder /app/stock-code /app/stock-code
COPY entrypoint /app/entrypoint
RUN chmod +x /app/entrypoint
ENTRYPOINT /app/entrypoint