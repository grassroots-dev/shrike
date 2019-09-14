
FROM golang:latest as goimage
ENV GO111MODULE=on
COPY Makefile .
COPY scripts ./scripts
RUN  apt-get update -y && \
    apt-get upgrade -y && \
    apt-get dist-upgrade -y && \
    apt-get -y autoremove && \
    apt-get clean
RUN apt-get install unzip \
    && rm -rf /var/lib/apt/lists/*
RUN make deps
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make protogen
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -v -o bin/shrike
FROM alpine:latest
WORKDIR /docker/bin
COPY --from=goimage /app/bin/shrike /docker/bin/
ENTRYPOINT ["/docker/bin/shrike"]
EXPOSE 1323
