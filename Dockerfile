# Build Stage
# First pull Golang image
FROM golang:1.21.5-bullseye as build
 
# Set environment variable
ENV APP_NAME freterapido-api
ENV CMD_PATH cmd/app/main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME

 
# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

 
# Run Stage
FROM alpine:3.14
 
# Set environment variable
ENV APP_NAME freterapido-api
 
# Copy only required data into this image
COPY --from=build /$APP_NAME .
COPY ./config.toml .
 
# Expose application port
EXPOSE 8080
 
# Start app
CMD ./$APP_NAME