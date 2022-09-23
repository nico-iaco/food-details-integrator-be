FROM golang:1.19.1-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /food-details-integrator-be

EXPOSE 8080

CMD [ "/food-details-integrator-be" ]