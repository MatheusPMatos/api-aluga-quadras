FROM golang:1.21.3-alpine3.17 as base
RUN apk update
WORKDIR /src/api
COPY go.mod go.sum ./
COPY . .
RUN go build -o apiRun ./cmd

FROM alpine:3.17 as binary

RUN apk --no-cache add tzdata
ENV TZ=America/Sao_Paulo

COPY --from=base /src/api/apiRun .

EXPOSE 8080
CMD [ "./apialugaquadras" ]
