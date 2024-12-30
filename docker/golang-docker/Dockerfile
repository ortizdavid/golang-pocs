FROM golang:latest 

# copy directory
WORKDIR /app

# copy files
COPY . .

# build or execute
RUN GOOS=linux go build -o golang-docker

# expose in a port
EXPOSE 8000

# command
CMD [ "./golang-docker" ]