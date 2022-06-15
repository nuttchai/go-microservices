# # FIRST IMAGE
# # base go image
# FROM golang:1.18-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# # CGO_ENABLED is enviroment variable 
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# # make sure the binary (brokerApp) is executable
# RUN chmod +x /app/brokerApp

# # SECOND IMAGE
# # build a reduced size docker image (only take the binary)
# FROM alpine:latest

# RUN mkdir /app

# # copy the binary from above image (builder) to the /app directory
# COPY --from=builder /app/brokerApp /app 

# CMD ["/app/brokerApp"]

FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app 

CMD ["/app/brokerApp"]