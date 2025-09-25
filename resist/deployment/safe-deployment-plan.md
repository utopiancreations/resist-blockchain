# 🛡️ Safe Resist Blockchain Deployment Plan

## Overview
This plan ensures your **family chat app** and **business CRM** remain untouched while adding the Resist blockchain to your server.

## 🔍 Phase 1: Discovery & Documentation (REQUIRED FIRST)

### Step 1: Run the Audit Script
```bash
# Copy this to your server and run:
chmod +x server-audit.sh
./server-audit.sh > server-audit-$(date +%Y%m%d).txt

# Review the output to identify:
# - What ports are in use
# - Where your apps are located
# - What services are running
```

### Step 2: Document Your Critical Apps
Create a file documenting your existing apps:

```bash
# Family Chat App
- Location: /path/to/chat/app
- Port: ????
- Database: ????
- Service name: ????
- How to start/stop: ????

# Business CRM
- Location: /path/to/crm
- Port: ????
- Database: ????
- Service name: ????
- How to start/stop: ????
```

## 🎯 Phase 2: Safe Backup Strategy

### Full Server Backup (CRITICAL)
```bash
# Create backup directory
sudo mkdir -p /backup/$(date +%Y%m%d)

# Backup critical directories
sudo rsync -av /var/www/ /backup/$(date +%Y%m%d)/www/
sudo rsync -av /opt/ /backup/$(date +%Y%m%d)/opt/
sudo rsync -av /home/ /backup/$(date +%Y%m%d)/home/
sudo rsync -av /etc/ /backup/$(date +%Y%m%d)/etc/

# Backup databases (adjust as needed)
# MySQL
mysqldump --all-databases > /backup/$(date +%Y%m%d)/mysql-all.sql

# PostgreSQL
sudo -u postgres pg_dumpall > /backup/$(date +%Y%m%d)/postgres-all.sql

# Create a restoration guide
cat > /backup/$(date +%Y%m%d)/RESTORE-INSTRUCTIONS.txt << 'EOF'
To restore if something goes wrong:
1. Stop all services
2. Restore files from this backup
3. Restart original services
4. Test family chat app and CRM
EOF
```

## 🐳 Phase 3: Isolated Docker Deployment

### Port Allocation Strategy
```
EXISTING APPS (to be preserved):
- Family Chat: Port ????  (identify first)
- Business CRM: Port ????  (identify first)
- SSH: 22
- HTTP/HTTPS: 80, 443 (may be used by existing apps)

RESIST BLOCKCHAIN (new ports):
- Blockchain P2P: 26656
- Blockchain RPC: 26657
- Blockchain API: 1317
- Mobile API: 3000 (or 3001 if 3000 is taken)
- IPFS: 4001, 5001
- Monitoring: 3002, 9091
```

### Docker-Only Installation
This keeps Resist completely isolated from your existing setup:

```bash
# Create resist user (separate from your existing apps)
sudo useradd -m -s /bin/bash resist-blockchain

# Create isolated directory
sudo mkdir -p /opt/resist-blockchain
sudo chown resist-blockchain:resist-blockchain /opt/resist-blockchain

# Install Docker (if not already installed)
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker resist-blockchain

# Switch to resist user for all blockchain operations
sudo su - resist-blockchain
```

## 🔧 Phase 4: Modified Docker Compose

Here's a **conflict-free** version of the Docker setup:

```yaml
# /opt/resist-blockchain/docker-compose.yml
version: '3.8'

services:
  resist-node:
    image: resist-blockchain:latest
    container_name: resist-node
    restart: unless-stopped
    ports:
      - "26656:26656"  # P2P (unique to blockchain)
      - "26657:26657"  # RPC (unique to blockchain)
      - "1317:1317"    # API (unique to blockchain)
    volumes:
      - ./data/blockchain:/root/.resist
    networks:
      - resist-net
    mem_limit: 6g  # Reduced to leave room for your apps

  mobile-api:
    build: ./mobile-api
    container_name: resist-mobile-api
    restart: unless-stopped
    ports:
      - "3001:3000"    # Use 3001 to avoid conflicts
    depends_on:
      - resist-node
    networks:
      - resist-net
    mem_limit: 1g

  ipfs-node:
    image: ipfs/kubo:latest
    container_name: resist-ipfs
    restart: unless-stopped
    ports:
      - "4001:4001"    # IPFS swarm
      - "5001:5001"    # IPFS API
      - "8081:8080"    # IPFS gateway (changed from 8080)
    volumes:
      - ./data/ipfs:/data/ipfs
    networks:
      - resist-net
    mem_limit: 2g

  # NO NGINX - Use existing web server or separate subdomain

networks:
  resist-net:
    driver: bridge
    ipam:
      config:
        - subnet: 172.20.0.0/16  # Unique subnet
```

## 🌐 Phase 5: Web Server Integration

### Option A: Subdomain Approach (RECOMMENDED)
If you own a domain, create a subdomain for Resist:
- Family Chat: `chat.yourdomain.com` (existing)
- Business CRM: `crm.yourdomain.com` (existing)
- Resist Blockchain: `resist.yourdomain.com` (new)

### Option B: Port-Based Access
- Family Chat: `yourserver:existing_port`
- Business CRM: `yourserver:existing_port`
- Resist Mobile API: `yourserver:3001`

### Option C: Path-Based Routing
Add to your existing Nginx config:
```nginx
# Add this to existing server block
location /resist/ {
    proxy_pass http://localhost:3001/;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
}
```

## ⚡ Phase 6: Testing & Validation

### Pre-Deployment Checklist
- [ ] Server audit completed and saved
- [ ] Family chat app and CRM documented
- [ ] Full backup created
- [ ] Port conflicts identified and resolved
- [ ] Docker installed without disrupting existing services

### Safe Deployment Process
1. **Test in isolation**: Start only the blockchain container first
2. **Verify no conflicts**: Check that existing apps still work
3. **Gradual rollout**: Add one service at a time
4. **Monitor resources**: Watch CPU/RAM usage with `htop`
5. **Validate everything**: Test all three systems (chat, CRM, blockchain)

### Rollback Plan
If anything goes wrong:
```bash
# Stop Resist services immediately
docker-compose down

# Restore from backup if needed
sudo systemctl stop resist-blockchain
# Restore files from /backup/$(date)/
sudo systemctl restart your-existing-services
```

## 🎯 Resource Allocation with Your Apps

With 32GB RAM, here's a safe allocation:
```
Family Chat App:     2-4GB   (existing)
Business CRM:        2-4GB   (existing)
System & Other:      4GB     (existing)
---
Available for Resist: 20-24GB (plenty!)

Resist Allocation:
- Blockchain Node:   6GB
- Mobile API:        1GB
- IPFS:             2GB
- Monitoring:       1GB
- Buffer:           10GB+
```

## 📞 Emergency Contacts & Recovery

### If Something Goes Wrong
1. **Stop Resist immediately**: `docker-compose down`
2. **Check your critical apps**: Test family chat and CRM
3. **Restore from backup** if needed
4. **Contact support** with the audit report

### Success Validation
- [ ] Family chat app still accessible
- [ ] Business CRM still functional
- [ ] Resist blockchain running
- [ ] All three systems coexisting peacefully
- [ ] Server performance stable

---

**Remember**: The beauty of Docker is isolation. Your existing apps will be completely unaffected by the Resist blockchain deployment when done correctly.