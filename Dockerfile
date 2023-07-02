# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files to the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Download and install wkhtmltopdf
RUN apt-get update && apt-get install -y wkhtmltopdf

# Install wgo - Live Reload
RUN go install github.com/bokwoon95/wgo@latest

# Expose port 8080 for the web server
EXPOSE 8080

# Define the command to run the Go application
CMD ["wgo", "run", "-verbose", "./cmd/consumer/main.go"]