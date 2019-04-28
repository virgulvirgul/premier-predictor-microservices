# Docker Parent Image with Node and Typescript
FROM sandrokeil/typescript:latest

# Create Directory for the Container
WORKDIR /app

# Copy the files we need to our new Directory
ADD . /app

# Expose the port outside of the container
EXPOSE 3006

# Grab dependencies and transpile src directory to dist
RUN npm install && tsc

# Start the server
ENTRYPOINT ["node", "dist/server.js"]