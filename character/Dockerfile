FROM alpine

# Copy code from local package files into container's workspace
COPY ./dist/bin/character /go/bin/character

# Execute the binary
ENTRYPOINT /go/bin/character

# Document that the service listens on port 8080
EXPOSE 8080
