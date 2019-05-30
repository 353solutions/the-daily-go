# Docker container for nlpd
# Run
#   docker run -p8080:8080 tdg/nlp


FROM golang:1.12.5-stretch as build

WORKDIR /nlp
COPY . .
ENV CGO_ENABLED=0
RUN go build ./cmd/nlpd

FROM debian:stretch-slim
COPY --from=build /nlp/nlp /usr/local/bin

ENV NLP_PORT 8080
CMD nlpd
