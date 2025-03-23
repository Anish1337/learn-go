FROM debian:stretch-slim

# Set the PORT environment variable (default: 8080)
ENV PORT=8080

# Copy the compiled binary into the container
ADD learn-go /bin/learn-go

# Expose the port to allow external access
EXPOSE 8080

# Run the application
CMD ["/bin/learn-go"]

