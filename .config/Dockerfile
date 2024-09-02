# Set default arguments
ARG grafana_version=latest
ARG grafana_image=grafana-enterprise

# Use specified Grafana base image
FROM grafana/${grafana_image}:${grafana_version}

# Additional build arguments
ARG development=false
ARG GO_VERSION=1.21.6
ARG GO_ARCH=amd64

# Environment variables
ENV DEV "${development}"

# Set development-friendly Grafana environment variables
ENV GF_AUTH_ANONYMOUS_ORG_ROLE "Admin"
ENV GF_AUTH_ANONYMOUS_ENABLED "true"
ENV GF_AUTH_BASIC_ENABLED "false"
ENV GF_DEFAULT_APP_MODE "development"

# Custom label to reflect the maintainer and purpose
LABEL maintainer="Cloud Rhino <cloudrhinoltd@gmail.com>" \
      description="InfinityPlus Grafana Datasource plugin development environment with Go and debugging tools" \
      version="${grafana_version}"

# Set working directory
ENV GF_PATHS_HOME="/usr/share/grafana"
WORKDIR $GF_PATHS_HOME

# Use root user for package installations
USER root

# Install necessary tools for development mode
RUN if [ "${development}" = "true" ]; then \
    if grep -i -q alpine /etc/issue; then \
        apk add --no-cache supervisor inotify-tools git; \
    elif grep -i -q ubuntu /etc/issue; then \
        DEBIAN_FRONTEND=noninteractive apt-get update && \
        apt-get install -y --no-install-recommends supervisor inotify-tools git && \
        rm -rf /var/lib/apt/lists/*; \
    else \
        echo 'ERROR: Unsupported base image' && /bin/false; \
    fi \
    fi

# Copy supervisor configuration files
COPY supervisord/supervisord.conf /etc/supervisor.d/supervisord.ini
COPY supervisord/supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Install Go for development
RUN if [ "${development}" = "true" ]; then \
    curl -O -L https://golang.org/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-${GO_ARCH}.tar.gz && \
    echo "export PATH=$PATH:/usr/local/go/bin:~/go/bin" >> ~/.bashrc && \
    rm -f go${GO_VERSION}.linux-${GO_ARCH}.tar.gz; \
    fi

# Install Delve for debugging
RUN if [ "${development}" = "true" ]; then \
    /usr/local/go/bin/go install github.com/go-delve/delve/cmd/dlv@latest; \
    fi

# Install Mage for plugin building and rebuilding
RUN if [ "${development}" = "true" ]; then \
    git clone https://github.com/magefile/mage; \
    cd mage; \
    export PATH=$PATH:/usr/local/go/bin; \
    go run bootstrap.go; \
    fi

# Inject livereload script into Grafana's index.html for development purposes
RUN if [ "${development}" = "true" ]; then \
    sed -i 's|</body>|<script src="http://localhost:35729/livereload.js"></script></body>|g' /usr/share/grafana/public/views/index.html; \
    fi

# Copy custom entrypoint script
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Set the entrypoint
ENTRYPOINT ["/entrypoint.sh"]
