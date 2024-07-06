FROM golang:alpine AS builder
RUN mkdir /server
ADD . /server/
WORKDIR /server
RUN GOOS=linux go build -o GophKeeper_server  cmd/server/main.go

FROM alpine:latest as GophKeeper
RUN mkdir /server
WORKDIR /server
ENV DATABASE_DSN="${DATABASE_DSN}"
ENV GRPC_ADDRESS="${GRPC_ADDRESS}"
ENV FILES="${FILES}"
COPY --from=builder /server/GophKeeper_server ./
CMD ["./GophKeeper_server"]