# Use the latest Node.js as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Install pnpm
RUN npm install -g pnpm

# Copy the package.json and pnpm-lock.yaml files to the working directory
COPY package.json  ./

# Install the dependencies
RUN pnpm install

# Copy the rest of the application code to the working directory
COPY . .

# Build the Nuxt.js application
RUN pnpm run build

# Expose the port that the application will run on
EXPOSE 3000

# Start the application
CMD ["pnpm", "run", "dev"]
