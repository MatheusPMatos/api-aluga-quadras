FROM golang:1.21.3-alpine3.17 as base
RUN apk update
WORKDIR /src/apialugaquadras
COPY go.mod go.sum ./
COPY . .
RUN go build -o API-ALUGA-QUADRAS ./cmd/api

FROM alpine:3.17 as binary

RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo

COPY --from=base /src/apialugaquadras/apialugaquadras .

EXPOSE 8080
CMD [ "./apialugaquadras" ]
