FROM golang:1.20-alpine

WORKDIR /app


# If go.mod and go.sum do not exist, initialize a new Go module
RUN  go mod init Company-Management

# Download dependencies (if go.mod or go.sum are available)
RUN go mod tidy

# Copy the entire project (everything else)
COPY . .

RUN go get

RUN go build -o /Company-Management

EXPOSE 30000

# Run the executable
CMD ["/Company-Management"]
