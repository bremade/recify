FROM node:14.8.0-stretch AS frontend-builder
WORKDIR /opt
COPY ./frontend .
RUN npm install
RUN npm run build

FROM golang:1.15.0-buster AS backend-builder
WORKDIR /opt
COPY . .
RUN go build -o recify_bin

FROM debian:buster
WORKDIR /opt
COPY --from=frontend-builder /opt ./frontend
COPY --from=backend-builder /opt/recify_bin ./recify_bin

ENV GIN_MODE=release
CMD ["/opt/recify_bin"]
