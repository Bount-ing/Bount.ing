# Stage 1: Build the application
FROM node:18 as builder

# Set the working directory inside the container
WORKDIR /app

# Install pnpm
RUN npm install -g pnpm

# Copy the package.json and pnpm-lock.yaml files to the working directory
COPY package.json pnpm-lock.yaml ./

# Install the dependencies
RUN pnpm install --frozen-lockfile

# Copy the rest of the application code to the working directory
COPY . .

# Build the Nuxt.js application
RUN pnpm run build

# Stage 2: Serve the built application
FROM nginx:alpine

# Copy built files from the previous stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Copy nginx configuration
COPY ./nginx.conf /etc/nginx/nginx.conf

# Expose the port that the application will run on
EXPOSE 3000

# Start nginx
CMD ["nginx", "-g", "daemon off;"]

