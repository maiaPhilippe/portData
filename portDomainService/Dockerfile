FROM golang:1.16-alpine

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/port/portDomainService/

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
RUN go build

# This container exposes port 9090 to the outside world
EXPOSE 9090

# Run the executable
CMD ["./portDomainService"]