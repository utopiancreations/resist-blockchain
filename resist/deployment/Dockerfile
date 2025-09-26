# Resist Blockchain Node Dockerfile - Mock Service for Testing
FROM alpine:latest

RUN apk --no-cache add ca-certificates bash curl jq

# Create resist user
RUN addgroup -g 1000 resist && \
    adduser -D -s /bin/bash -u 1000 -G resist resist

WORKDIR /home/resist

# Create a simple mock blockchain service script
RUN cat > mock-resistd.sh << 'EOF'
#!/bin/bash
echo "Starting Mock Resist Blockchain Node..."
echo "Chain ID: resist-mainnet-1"
echo "Node Home: /home/resist/.resist"

# Create necessary directories
mkdir -p /home/resist/.resist/config
mkdir -p /home/resist/.resist/data

# Mock health endpoint server
while true; do
    echo "$(date): Resist blockchain node is running on ports 26656 (P2P), 26657 (RPC), 1317 (API)"
    sleep 30
done
EOF

RUN chmod +x mock-resistd.sh && chown resist:resist mock-resistd.sh

# Expose blockchain ports
EXPOSE 26656 26657 1317

# Switch to resist user
USER resist

# Start the mock service
CMD ["./mock-resistd.sh"]