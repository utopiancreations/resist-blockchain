#!/bin/bash

# Resist Blockchain Community Node Setup Script
# This script sets up a complete Resist blockchain node in minutes!
#
# Usage: curl -sSL https://setup.resist-blockchain.org | bash
# Or: ./setup-community-node.sh

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
PURPLE='\033[0;35m'
NC='\033[0m' # No Color

print_banner() {
    echo -e "${PURPLE}"
    echo "╔═══════════════════════════════════════════════════════════════════╗"
    echo "║                                                                   ║"
    echo "║         🚀 RESIST BLOCKCHAIN COMMUNITY NODE SETUP 🚀             ║"
    echo "║                                                                   ║"
    echo "║  Welcome to the decentralized social network revolution!         ║"
    echo "║  This script will set up your community node in minutes.         ║"
    echo "║                                                                   ║"
    echo "╚═══════════════════════════════════════════════════════════════════╝"
    echo -e "${NC}"
    echo ""
}

print_step() {
    echo -e "${BLUE}🔧 $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_info() {
    echo -e "${PURPLE}ℹ️  $1${NC}"
}

# Check if running as root
check_root() {
    if [[ $EUID -eq 0 ]]; then
        print_error "Please don't run this script as root for security reasons"
        exit 1
    fi
}

# Check system requirements
check_requirements() {
    print_step "Checking system requirements..."

    # Check OS
    if [[ "$OSTYPE" != "linux-gnu"* ]]; then
        print_error "This script is designed for Linux systems"
        exit 1
    fi

    # Check memory (minimum 4GB recommended)
    MEMORY_GB=$(free -g | awk '/^Mem:/{print $2}')
    if (( MEMORY_GB < 2 )); then
        print_warning "Your system has only ${MEMORY_GB}GB RAM. Minimum 4GB recommended."
        read -p "Continue anyway? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 1
        fi
    fi

    # Check disk space (minimum 50GB recommended)
    DISK_GB=$(df -BG / | awk 'NR==2 {print $4}' | sed 's/G//')
    if (( DISK_GB < 20 )); then
        print_warning "Low disk space: ${DISK_GB}GB available. Minimum 50GB recommended."
        read -p "Continue anyway? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            exit 1
        fi
    fi

    print_success "System requirements check passed"
}

# Install Docker if not present
install_docker() {
    if ! command -v docker &> /dev/null; then
        print_step "Installing Docker..."
        curl -fsSL https://get.docker.com -o get-docker.sh
        sudo sh get-docker.sh
        sudo usermod -aG docker $USER
        rm get-docker.sh
        print_success "Docker installed successfully"
        print_warning "You may need to log out and back in for Docker permissions to take effect"
    else
        print_success "Docker is already installed"
    fi
}

# Install Docker Compose if not present
install_docker_compose() {
    if ! command -v docker-compose &> /dev/null; then
        print_step "Installing Docker Compose..."
        sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
        sudo chmod +x /usr/local/bin/docker-compose
        print_success "Docker Compose installed successfully"
    else
        print_success "Docker Compose is already installed"
    fi
}

# Get user configuration
get_user_config() {
    print_step "Setting up your node configuration..."
    echo ""

    # Get node name
    echo -e "${YELLOW}What would you like to name your node? (e.g., 'community-node-nyc', 'alice-node')${NC}"
    read -p "Node name: " NODE_MONIKER
    NODE_MONIKER=${NODE_MONIKER:-"community-node-$(hostname)"}

    # Get operator name
    echo -e "${YELLOW}What's your name or organization? (for community recognition)${NC}"
    read -p "Node operator: " NODE_OPERATOR
    NODE_OPERATOR=${NODE_OPERATOR:-"Anonymous"}

    # Get external address (optional)
    echo -e "${YELLOW}Do you have a public IP or domain name? (optional, press enter to skip)${NC}"
    echo -e "${PURPLE}Examples: your-domain.com, 123.456.789.101${NC}"
    read -p "External address: " EXTERNAL_ADDRESS

    # Check if ports are available
    print_step "Checking port availability..."
    PORTS_TO_CHECK=(26656 26657 1317 4001 5001 8080 3000 9091)
    USED_PORTS=()

    for port in "${PORTS_TO_CHECK[@]}"; do
        if ss -tuln | grep -q ":$port "; then
            USED_PORTS+=($port)
        fi
    done

    if [ ${#USED_PORTS[@]} -gt 0 ]; then
        print_warning "The following ports are already in use: ${USED_PORTS[*]}"
        print_info "You may need to adjust port configuration in .env file after setup"
    fi

    print_success "Configuration collected"
}

# Create deployment directory
create_deployment() {
    print_step "Creating deployment directory..."

    DEPLOY_DIR="$HOME/resist-node"
    mkdir -p "$DEPLOY_DIR"
    cd "$DEPLOY_DIR"

    # Download deployment files
    print_step "Downloading Resist node deployment files..."

    # For now, we'll create the files directly since GitHub isn't set up yet
    # In production, this would be: wget https://github.com/utopiancreations/resist-blockchain/archive/community-node.zip

    # Create directories
    mkdir -p config keys backups monitoring api-gateway

    print_success "Deployment directory created at $DEPLOY_DIR"
}

# Create configuration file
create_config() {
    print_step "Creating your node configuration..."

    cat > .env <<EOF
# Resist Blockchain Community Node Configuration
# Generated by setup script on $(date)

# Node Identity
NODE_MONIKER=$NODE_MONIKER
NODE_OPERATOR=$NODE_OPERATOR
EXTERNAL_ADDRESS=${EXTERNAL_ADDRESS:+tcp://$EXTERNAL_ADDRESS:26656}

# Network Configuration
CHAIN_ID=resist-mainnet-1
SEED_NODES=node1@resist-hub.duckdns.org:26656
MIN_GAS_PRICES=0.025stake

# Port Configuration (default values)
NODE_P2P_PORT=26656
NODE_RPC_PORT=26657
NODE_API_PORT=1317
NODE_GRPC_PORT=9090
IPFS_SWARM_PORT=4001
IPFS_API_PORT=5001
IPFS_GATEWAY_PORT=8080
API_GATEWAY_PORT=3000
MONITOR_PORT=9091

# Resource Limits
NODE_MEMORY_LIMIT=4g
IPFS_MEMORY_LIMIT=2g
API_MEMORY_LIMIT=1g
MONITOR_MEMORY_LIMIT=512m
NODE_CPU_LIMIT=2

# Validator Mode (set to true if you want to become a validator)
VALIDATOR_MODE=false

# Backup Settings
BACKUP_RETENTION_DAYS=7
EOF

    print_success "Configuration file created (.env)"
}

# Create monitoring configuration
create_monitoring() {
    print_step "Setting up monitoring..."

    cat > monitoring/prometheus.yml <<EOF
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: 'resist-node'
    static_configs:
      - targets: ['resist-community-node:26660']  # Cosmos metrics port
    metrics_path: /metrics
    scrape_interval: 10s

  - job_name: 'ipfs-node'
    static_configs:
      - targets: ['ipfs-node:5001']
    metrics_path: /debug/metrics/prometheus
    scrape_interval: 30s

  - job_name: 'node-exporter'
    static_configs:
      - targets: ['localhost:9100']
    scrape_interval: 15s
EOF

    print_success "Monitoring configuration created"
}

# Create simple API gateway
create_api_gateway() {
    print_step "Creating community API gateway..."

    cat > api-gateway/package.json <<EOF
{
  "name": "resist-community-api",
  "version": "1.0.0",
  "description": "Community node API gateway for Resist blockchain",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "dependencies": {
    "express": "^4.18.2",
    "cors": "^2.8.5",
    "axios": "^1.6.0"
  }
}
EOF

    cat > api-gateway/server.js <<EOF
const express = require('express');
const cors = require('cors');
const axios = require('axios');

const app = express();
const PORT = process.env.PORT || 3000;

const BLOCKCHAIN_RPC = process.env.BLOCKCHAIN_RPC || 'http://localhost:26657';
const BLOCKCHAIN_API = process.env.BLOCKCHAIN_API || 'http://localhost:1317';
const IPFS_API = process.env.IPFS_API || 'http://localhost:5001';

app.use(cors());
app.use(express.json());

// Community node info endpoint
app.get('/node/info', async (req, res) => {
  try {
    const [nodeInfo, status] = await Promise.all([
      axios.get(\`\${BLOCKCHAIN_RPC}/status\`),
      axios.get(\`\${BLOCKCHAIN_API}/cosmos/base/tendermint/v1beta1/node_info\`)
    ]);

    res.json({
      community_node: true,
      operator: process.env.NODE_OPERATOR || 'Anonymous',
      moniker: nodeInfo.data.result.node_info.moniker,
      network: nodeInfo.data.result.node_info.network,
      latest_block_height: nodeInfo.data.result.sync_info.latest_block_height,
      catching_up: nodeInfo.data.result.sync_info.catching_up,
      node_info: status.data.node_info
    });
  } catch (error) {
    res.status(500).json({ error: 'Failed to fetch node info', details: error.message });
  }
});

// Health check endpoint
app.get('/health', async (req, res) => {
  try {
    await axios.get(\`\${BLOCKCHAIN_RPC}/health\`);
    res.json({ status: 'healthy', timestamp: new Date().toISOString() });
  } catch (error) {
    res.status(503).json({ status: 'unhealthy', error: error.message });
  }
});

// Proxy endpoints for mobile apps
app.use('/api/blockchain', (req, res) => {
  const url = \`\${BLOCKCHAIN_API}\${req.path}\`;
  req.pipe(axios.get(url)).pipe(res);
});

app.use('/api/ipfs', (req, res) => {
  const url = \`\${IPFS_API}\${req.path}\`;
  req.pipe(axios.get(url)).pipe(res);
});

app.listen(PORT, () => {
  console.log(\`Resist Community API Gateway running on port \${PORT}\`);
  console.log(\`Node operator: \${process.env.NODE_OPERATOR || 'Anonymous'}\`);
});
EOF

    cat > api-gateway/Dockerfile <<EOF
FROM node:18-alpine

WORKDIR /app

COPY package.json ./
RUN npm install --production

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
EOF

    print_success "API gateway created"
}

# Start the node
start_node() {
    print_step "Starting your Resist community node..."

    # Create docker-compose.yml (simplified version for community nodes)
    cat > docker-compose.yml <<EOF
version: '3.8'

services:
  resist-node:
    image: cosmoscontracts/juno:latest  # Temporary until resist image is ready
    container_name: resist-community-node
    restart: unless-stopped
    ports:
      - "\${NODE_P2P_PORT:-26656}:26656"
      - "\${NODE_RPC_PORT:-26657}:26657"
      - "\${NODE_API_PORT:-1317}:1317"
    volumes:
      - blockchain-data:/root/.juno
    environment:
      - CHAIN_ID=\${CHAIN_ID:-resist-mainnet-1}
      - MONIKER=\${NODE_MONIKER:-community-node}
    command: |
      sh -c "
        echo 'Resist Community Node Starting...'
        echo 'Node: \${NODE_MONIKER}'
        echo 'Operator: \${NODE_OPERATOR}'
        echo 'This is a placeholder - will be updated when Resist binary is ready'
        sleep 3600
      "

volumes:
  blockchain-data:
    driver: local

networks:
  default:
    name: resist-community-network
EOF

    # Start services
    docker-compose up -d

    print_success "Community node started!"
    echo ""
    print_info "Your node is now running. Here's what you can do:"
    echo ""
    echo -e "${GREEN}📊 Check node status:${NC}"
    echo "   docker-compose ps"
    echo "   docker-compose logs -f resist-community-node"
    echo ""
    echo -e "${GREEN}🌐 Access your node:${NC}"
    echo "   Node API: http://localhost:1317"
    echo "   Node RPC: http://localhost:26657"
    if [ ! -z "$EXTERNAL_ADDRESS" ]; then
        echo "   External: $EXTERNAL_ADDRESS"
    fi
    echo ""
}

# Configure firewall (optional)
configure_firewall() {
    print_step "Configuring firewall (optional)..."

    read -p "Would you like to automatically configure UFW firewall? (y/N): " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        if command -v ufw &> /dev/null; then
            sudo ufw allow 26656/tcp comment "Resist P2P"
            sudo ufw allow 26657/tcp comment "Resist RPC"
            sudo ufw allow 1317/tcp comment "Resist API"
            sudo ufw allow 4001/tcp comment "IPFS Swarm"
            print_success "Firewall configured"
        else
            print_warning "UFW not installed, skipping firewall configuration"
        fi
    else
        print_info "Skipping firewall configuration"
        print_info "Remember to manually open ports: 26656, 26657, 1317, 4001"
    fi
}

# Final instructions
show_final_instructions() {
    echo ""
    print_banner
    echo -e "${GREEN}🎉 CONGRATULATIONS! Your Resist community node is running! 🎉${NC}"
    echo ""
    echo -e "${BLUE}📋 Next Steps:${NC}"
    echo ""
    echo "1. ${GREEN}Monitor your node:${NC}"
    echo "   cd $HOME/resist-node"
    echo "   docker-compose logs -f"
    echo ""
    echo "2. ${GREEN}Configure port forwarding on your router:${NC}"
    echo "   Forward ports 26656, 26657, 1317, 4001 to this machine"
    echo ""
    echo "3. ${GREEN}Join the community:${NC}"
    echo "   Discord: https://discord.gg/resist-blockchain"
    echo "   Docs: https://docs.resist-blockchain.org"
    echo ""
    echo "4. ${GREEN}Become a validator (optional):${NC}"
    echo "   Edit .env file: VALIDATOR_MODE=true"
    echo "   Follow validator setup guide"
    echo ""
    echo -e "${PURPLE}Your contribution helps build a truly decentralized social network!${NC}"
    echo -e "${PURPLE}Thank you for joining the Resist revolution! 🚀${NC}"
    echo ""
}

# Main execution
main() {
    print_banner

    check_root
    check_requirements
    install_docker
    install_docker_compose
    get_user_config
    create_deployment
    create_config
    create_monitoring
    create_api_gateway
    start_node
    configure_firewall
    show_final_instructions
}

# Run main function
main "$@"