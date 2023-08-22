FROM golang:latest

# Set the Current Working Directory inside the container
RUN mkdir -p /app
WORKDIR /app
COPY . /app

# Download all the dependencies
RUN go get -d -v ./...
RUN go build -o main .

# This container exposes port 8000 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./main"]
