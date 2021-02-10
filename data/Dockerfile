FROM golang:1.15-buster as build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=unknown
RUN go build -ldflags="-X main.version=${VERSION}" -o nlpd ./cmd/nlpd

FROM debian:buster-slim
RUN mkdir /app
COPY --from=build /app/nlpd /app/nlpd

# Don't run as root user
RUN useradd nlp
USER nlp
ENV NLPD_ADDR 8080
CMD /app/nlpd
