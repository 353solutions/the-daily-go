# Docker container for nlpd
# Run
#   docker run -p8080:8080 tdg/nlp


FROM golang:1.12.1-stretch as build

WORKDIR /nlp
COPY . .
ARG NLP_VERSION=unknown
ENV CGO_ENABLED=0
RUN go build -ldflags="-X main.Version=${NLP_VERSION}" ./cmd/nlpd
RUN cp nlpd /usr/local/bin

FROM debian:stretch-slim
COPY --from=build /usr/local/bin/nlp /usr/local/bin

ENV NLP_PORT 8080
CMD nlpd
