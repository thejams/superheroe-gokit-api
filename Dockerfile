# Start from golang base image
FROM golang:alpine

# Install make to alpine image
RUN apk update && apk add --no-cache git && apk add --no-cache make

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build proyect
RUN make mod && make build

CMD ["./build/bin/superheroe-gokit-api"]
