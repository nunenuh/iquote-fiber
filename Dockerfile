FROM bitnami/minideb:bookworm

RUN apt-get update && apt-get install iputils-ping net-tools curl libc6 wget nano iproute2 -y

# Create a new user 'nonroot'
RUN useradd -r -u 1001 nonroot

# Create a new directory for the application and change its owner to 'nonroot'
RUN mkdir /app
RUN chown 1001:1001 /app

# Switch to the new user
USER 1001

# Set the new directory as the working directory
WORKDIR /app

# Change the ownership of the application directory to the new user
COPY --chown=1001:1001 ./bin/api /app/

# Run the application
CMD ["./api"]