FROM golang:1.14-alpine as builder

WORKDIR /go/src/web
COPY ./go/src/web/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -extldflags "-static"' -o main
RUN ls

FROM scratch
COPY --from=builder  /go/src/web .

CMD [ "./main" ]