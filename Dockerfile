FROM --platform=linux/amd64 golang:1.19-alpine as build-stage

WORKDIR /app

COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY *.go /app/

RUN go mod download
RUN go build -o /bin/entrypoint .

FROM --platform=linux/amd64 alpine:latest 

COPY --from=build-stage /bin/entrypoint /bin/entrypoint

EXPOSE 8080
CMD ["/bin/entrypoint"]
