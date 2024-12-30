FROM golang:latest

WORKDIR /app

# Copy dependencies
COPY go.mod go.sum ./ 

# install dependencies
RUN go mod download

# copy all files
COPY . .

# build application
RUN GOOS=linux go build -o go-fiber

# http port
EXPOSE 9000

# command
CMD [ "./go-fiber" ]