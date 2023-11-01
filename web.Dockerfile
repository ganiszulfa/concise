FROM node:18-buster-slim

# install dependencies
RUN mkdir -p /app
WORKDIR /app
COPY ./web/package.json /app
COPY ./web/package-lock.json /app
RUN echo $(npm --version)
RUN npm ci

# Copy all local files into the image.
COPY ./web/ /app

# Generate css to ./src/output.css
RUN npm run css

# Build the web to ./build
RUN npm run build

ENV SVELTE_MODE=release

EXPOSE 3000
CMD ["node", "./build"]