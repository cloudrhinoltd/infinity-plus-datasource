# Dockerfile

# Use the official Grafana image as a base
FROM grafana/grafana-enterprise:${GF_VERSION:-main}

# Switch to root user to install dependencies
USER root

# Install necessary dependencies including build tools
RUN apk update && \
    apk add --no-cache curl bash openssl ca-certificates tar gzip python3 py3-pip \
    gcc musl-dev linux-headers python3-dev

# Set up Python virtual environment and install Azure CLI
RUN python3 -m venv /opt/azcli-venv && \
    . /opt/azcli-venv/bin/activate && \
    pip install --upgrade pip && \
    pip install azure-cli && \
    deactivate

# Add virtual environment to PATH for all users
ENV PATH="/opt/azcli-venv/bin:$PATH"

# Revert back to Grafana default user for security purposes
USER grafana

# Continue with the base image's entrypoint
ENTRYPOINT [ "/run.sh" ]
