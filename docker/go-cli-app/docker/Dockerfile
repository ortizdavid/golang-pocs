FROM golang:latest 

WORKDIR /app

COPY go.mod ./

COPY . .

RUN GOOS=linux go build -o go-cli-app .

CMD [ "./go-cli-app" ]

