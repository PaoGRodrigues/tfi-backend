FROM golang:1.23-bullseye

RUN mkdir /build

WORKDIR /build

COPY ./app ./app

COPY go.mod go.sum ./

RUN mkdir /backend
RUN go build -o /backend/main ./app/

WORKDIR /backend
ENV GIN_MODE=release

RUN rm -fr /build

CMD ["/backend/main"]