FROM golang:1.19.1-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /food-details-integrator-be

EXPOSE 8080

FROM alpine

COPY --from=builder /food-details-integrator-be .

CMD [ "/food-details-integrator-be" ]