FROM ubuntu

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
ADD . /app

# Install any needed packages specified in requirements.txt
RUN apt-get update && apt-get install -y \
  golang-go \
  git

# Remove the previously built executable
RUN rm -rf ./go

RUN go build

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Define environment variable
ENV LOGGING_LEVEL debug

CMD ["./app"]
