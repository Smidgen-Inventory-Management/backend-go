FROM golang:1.22 AS build
WORKDIR /go/
COPY go ./go
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o smidgen-backgend .

FROM scratch AS runtime
COPY --from=build /go/ ./
EXPOSE 8050/tcp
ENTRYPOINT ["./smidgen-backgend"]
