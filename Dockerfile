# syntax=docker/dockerfile:1

# Fetch GO
FROM golang:1.23rc2

# Create sub directory
WORKDIR /app

# Copy dependency files into container
# COPY go.mod ./
# COPY go.sum ./

# Copy everything from project into container
COPY . ./

# Download all dependencies from go mod
RUN go mod download

# Build binary
RUN go build -o /minitwit

EXPOSE 8080

CMD [ "/minitwit" ]