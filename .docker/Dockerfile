# Start with an up-to-date Hugo Extended image (supports SASS/SCSS layouts)
FROM hugomods/hugo:ext-debian

# Set the working directory inside the container
WORKDIR /src

# Copy all your new Hugo project files into the container
COPY . .

# Expose Hugo's default local development port
EXPOSE 1313

# Start the Hugo local development server automatically
CMD ["hugo", "server", "--bind", "0.0.0.0"]