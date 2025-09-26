# 🚀 Resist Blockchain Community Node

Welcome to the **Resist Network** - a decentralized social media blockchain built for freedom, privacy, and community ownership!

## 🌍 What is a Community Node?

A Community Node is your contribution to the Resist decentralized social network. By running a node, you:

- **Help secure the network** through distributed consensus
- **Earn token rewards** for providing resources and uptime
- **Support content distribution** via integrated IPFS
- **Enable mobile users** to connect through your node
- **Participate in governance** of the network's future

## ⚡ Quick Start (5 Minutes)

### Option 1: Automated Setup Script (Recommended)
```bash
curl -sSL https://setup.resist-blockchain.org/install.sh | bash
```

### Option 2: Manual Setup
```bash
git clone https://github.com/utopiancreations/resist-blockchain.git
cd resist-blockchain/deployment/community-node
cp .env.template .env
# Edit .env with your configuration
docker-compose up -d
```

## 📋 Requirements

### Minimum Requirements
- **OS:** Linux (Ubuntu 20.04+ recommended)
- **RAM:** 4GB (8GB+ recommended)
- **Storage:** 50GB SSD (100GB+ recommended)
- **Network:** Broadband internet with stable connection
- **Ports:** Ability to open ports 26656, 26657, 1317, 4001

### Recommended Requirements
- **RAM:** 16GB+ for validator nodes
- **Storage:** 200GB+ SSD for long-term operation
- **CPU:** 4+ cores for optimal performance
- **Network:** Unlimited bandwidth or high data cap

## 🔧 Configuration Options

### Basic Configuration (`.env` file)
```bash
# Your node's unique name
NODE_MONIKER=my-awesome-node

# Your name or organization
NODE_OPERATOR=YourName

# Your public IP or domain (if you have one)
EXTERNAL_ADDRESS=your-domain.com

# Validator mode (requires additional setup)
VALIDATOR_MODE=false
```

### Port Configuration
If you have port conflicts, adjust these in your `.env`:
```bash
NODE_P2P_PORT=26656        # Blockchain P2P
NODE_RPC_PORT=26657        # Blockchain RPC API
NODE_API_PORT=1317         # REST API
IPFS_SWARM_PORT=4001       # IPFS peer connections
```

## 🛠️ Common Operations

### Start Your Node
```bash
docker-compose up -d
```

### Check Node Status
```bash
# View all services
docker-compose ps

# Check node logs
docker-compose logs -f resist-community-node

# Check node sync status
curl http://localhost:26657/status
```

### Stop Your Node
```bash
docker-compose down
```

### Update Your Node
```bash
docker-compose pull
docker-compose up -d
```

### Backup Your Node
```bash
# Manual backup
docker-compose run --rm backup-service

# View backups
ls -la backups/
```

## 🏆 Becoming a Validator

Validators help secure the network and earn higher rewards. To become a validator:

1. **Ensure your node is synced:**
   ```bash
   curl http://localhost:26657/status | grep catching_up
   # Should show "catching_up": false
   ```

2. **Set validator mode:**
   ```bash
   # Edit .env file
   VALIDATOR_MODE=true

   # Restart services
   docker-compose down && docker-compose up -d
   ```

3. **Create validator transaction:**
   ```bash
   # Generate validator keys
   docker-compose exec resist-community-node resistd keys add validator

   # Create validator
   docker-compose exec resist-community-node resistd tx staking create-validator \
     --amount=1000000stake \
     --pubkey=$(resistd tendermint show-validator) \
     --moniker="$NODE_MONIKER" \
     --commission-rate="0.10" \
     --commission-max-rate="0.20" \
     --commission-max-change-rate="0.01" \
     --min-self-delegation="1" \
     --from=validator \
     --chain-id=resist-mainnet-1
   ```

## 💰 Node Rewards & Economics

### Reward Sources
- **Block Rewards:** Earn tokens for validating transactions
- **Transaction Fees:** Share of network transaction fees
- **Content Hosting:** Rewards for IPFS content distribution
- **Network Uptime:** Bonus rewards for high availability
- **Community Participation:** Governance and voting rewards

### Resource Marketplace
Your node automatically participates in the resource marketplace:
- **Storage:** Earn tokens for hosting content via IPFS
- **Bandwidth:** Rewards for serving mobile users
- **Compute:** Processing rewards for complex operations
- **Geographic Value:** Higher rewards for underserved regions

## 🔐 Security Best Practices

### Server Security
- **Firewall:** Only open required ports (26656, 26657, 1317, 4001)
- **SSH:** Use key-based authentication, disable password login
- **Updates:** Keep your system and Docker updated
- **Monitoring:** Set up alerts for node downtime

### Key Management
- **Backup keys:** Store validator keys securely offline
- **Use hardware wallets** for large stake amounts
- **Separate keys:** Use different keys for different purposes
- **Regular rotation:** Change keys periodically

### Network Security
- **VPN:** Consider running behind a VPN for additional privacy
- **DDoS Protection:** Use services like Cloudflare for public nodes
- **Rate Limiting:** Configure appropriate API rate limits

## 📊 Monitoring & Troubleshooting

### Built-in Monitoring
Access your node's monitoring dashboard:
```bash
# Prometheus metrics
http://localhost:9091

# Node status API
http://localhost:1317/node/info

# Health check
curl http://localhost:3000/health
```

### Common Issues

#### Node Won't Sync
1. Check internet connection and firewall
2. Verify seed nodes in configuration
3. Check if ports are accessible from outside

#### High Resource Usage
1. Adjust memory limits in `.env`
2. Clean old blockchain data: `docker-compose exec resist-community-node resistd unsafe-reset-all`
3. Monitor disk usage and clean old backups

#### Connection Issues
1. Verify `EXTERNAL_ADDRESS` is correct
2. Check router port forwarding
3. Test ports with external tools

### Log Analysis
```bash
# View recent logs
docker-compose logs --tail=100 resist-community-node

# Follow logs in real-time
docker-compose logs -f

# Search for specific errors
docker-compose logs resist-community-node | grep ERROR
```

## 🌐 Network Integration

### Router Configuration
Forward these ports to your node's local IP:
- `26656` → Blockchain P2P
- `26657` → Blockchain RPC
- `1317` → REST API
- `4001` → IPFS Swarm

### Dynamic DNS Setup
For home networks with changing IPs:
1. Set up dynamic DNS (DuckDNS, No-IP, etc.)
2. Configure your router to update DNS automatically
3. Use your domain name in `EXTERNAL_ADDRESS`

### Mobile App Integration
Your node automatically serves mobile users:
- **Lite clients** sync through your RPC endpoint
- **Content distribution** via your IPFS node
- **API gateway** provides optimized mobile endpoints

## 🤝 Community & Support

### Get Help
- **Documentation:** https://docs.resist-blockchain.org
- **Discord:** https://discord.gg/resist-blockchain
- **GitHub Issues:** https://github.com/utopiancreations/resist-blockchain/issues
- **Community Forum:** https://forum.resist-blockchain.org

### Contribute
- **Run a node** and help secure the network
- **Report bugs** and suggest improvements
- **Share resources** and help other node operators
- **Participate in governance** voting and proposals

### Node Operator Recognition
- **Leaderboards:** Top nodes by uptime, performance, contribution
- **Community badges:** Recognition for long-term operators
- **Governance power:** Node operators get enhanced voting rights
- **Economic benefits:** Higher rewards for established operators

## 🚀 What's Next?

1. **Start your node** and let it sync to the network
2. **Join our Discord** and introduce yourself
3. **Configure monitoring** to track your node's performance
4. **Consider becoming a validator** for higher rewards
5. **Help grow the network** by encouraging others to run nodes

## 📈 Roadmap

### Short Term (1-3 months)
- Enhanced mobile app integration
- Automated node updates
- Advanced monitoring dashboards
- Geographic reward multipliers

### Medium Term (3-6 months)
- Cross-chain bridges and integrations
- Enhanced governance features
- Mobile node capabilities
- Enterprise APIs

### Long Term (6+ months)
- Sharding for massive scale
- Advanced privacy features
- AI-powered content curation
- Global mesh networking

---

**Welcome to the Resist revolution!** 🚀

Your node helps build a truly decentralized, censorship-resistant social network owned by its users. Every node makes the network stronger, more resilient, and more valuable for everyone.

*Thank you for contributing to digital freedom and decentralization!*