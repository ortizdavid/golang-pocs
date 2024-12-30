FROM golang:latest

WORKDIR /app 

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux go build -o go-gin

EXPOSE 7000

CMD [ "./go-gin" ]