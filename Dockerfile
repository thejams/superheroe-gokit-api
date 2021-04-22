# Start from golang base image
FROM golang:alpine

# Build proyect
RUN apk update && apk add --no-cache git && apk add --no-cache make

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build proyect
RUN make mod && make build

# Expose port 8080 in the container
EXPOSE 8080

CMD ["./build/bin/superheroe-gokit-api"]
