FROM golang:1.21.1
WORKDIR /app
COPY . .
EXPOSE 50051
RUN go build -o ./bin/roller_server ./roller/roller_server
# Command to run the application
CMD ["./bin/roller_server"]