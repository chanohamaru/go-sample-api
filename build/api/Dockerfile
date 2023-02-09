FROM golang:1.17.1-alpine as builder

WORKDIR /builder
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/api

FROM alpine
# ビルドした内容をbuilderの/build/apiからコピー
COPY --from=builder /builder/api /opt/app/
ENTRYPOINT ["/opt/app/api"]