FROM ubuntu

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
ADD . /app

# Install the pre-reqs
RUN apt-get update && apt-get install -y \
  curl \
  python-software-properties

# Add a new source for node 8 since Ubuntu is super legacy here
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -

RUN apt-get update && apt-get install -y \
  nodejs

# Remove the previously built executable
RUN rm -rf ./dist

# Remove node_nodules
RUN rm -rf node_modules

# Install dependencies
RUN npm install

# Build
RUN npm run build

# Make port 8080 available to the world outside this container
EXPOSE 8080

# Define environment variable
ENV LOGGING_LEVEL debug

CMD ["nodejs", "./dist/main.js"]
