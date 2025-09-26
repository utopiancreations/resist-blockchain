#!/bin/bash

# Server Audit Script - Identify existing services safely
# This script will help you understand what's currently running

set -e

echo "🔍 RESIST BLOCKCHAIN - SERVER AUDIT REPORT"
echo "=========================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_section() {
    echo -e "${BLUE}📋 $1${NC}"
    echo "----------------------------------------"
}

print_important() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_safe() {
    echo -e "${GREEN}✅ $1${NC}"
}

# System Overview
print_section "SYSTEM OVERVIEW"
echo "Hostname: $(hostname)"
echo "OS: $(lsb_release -d 2>/dev/null | cut -f2 || cat /etc/os-release | grep PRETTY_NAME | cut -d'"' -f2)"
echo "Uptime: $(uptime -p)"
echo "CPU: $(nproc) cores - $(lscpu | grep 'Model name' | cut -d':' -f2 | sed 's/^ *//')"
echo "RAM: $(free -h | awk '/^Mem:/ {print $2}') total, $(free -h | awk '/^Mem:/ {print $3}') used"
echo "Disk: $(df -h / | awk 'NR==2 {print $2}') total, $(df -h / | awk 'NR==2 {print $3}') used"
echo ""

# Running Services
print_section "RUNNING SERVICES"
echo "Active systemd services:"
systemctl list-units --state=running --type=service --no-pager | grep -E "(chat|crm|web|http|nginx|apache|mysql|postgres|mongo|redis|docker)" || echo "No obvious web services found in systemd"
echo ""

# Docker Containers (if Docker is installed)
if command -v docker &> /dev/null; then
    print_section "DOCKER CONTAINERS"
    if docker ps -q &> /dev/null; then
        echo "Running containers:"
        docker ps --format "table {{.Names}}\t{{.Image}}\t{{.Ports}}\t{{.Status}}" 2>/dev/null || echo "No running containers"
        echo ""
        echo "All containers (including stopped):"
        docker ps -a --format "table {{.Names}}\t{{.Image}}\t{{.Status}}" 2>/dev/null || echo "No containers found"
    else
        echo "Docker is installed but not accessible (may need sudo)"
    fi
else
    echo "Docker not installed"
fi
echo ""

# Network Ports
print_section "NETWORK PORTS IN USE"
echo "Listening services:"
netstat -tlnp 2>/dev/null | grep LISTEN | head -20 || ss -tlnp | grep LISTEN | head -20
echo ""

# Web Servers
print_section "WEB SERVERS"
if pgrep -x "nginx" > /dev/null; then
    print_important "Nginx is running"
    echo "Nginx config files:"
    find /etc/nginx -name "*.conf" -o -name "*sites-*" 2>/dev/null | head -10
    echo "Active sites:"
    ls -la /etc/nginx/sites-enabled/ 2>/dev/null || echo "Sites-enabled directory not found"
fi

if pgrep -x "apache2" > /dev/null || pgrep -x "httpd" > /dev/null; then
    print_important "Apache is running"
    echo "Apache config:"
    ls -la /etc/apache2/sites-enabled/ 2>/dev/null || ls -la /etc/httpd/conf.d/ 2>/dev/null || echo "Apache config not in standard location"
fi
echo ""

# Databases
print_section "DATABASES"
for db in mysql postgresql mongodb redis; do
    if pgrep -x "$db" > /dev/null || pgrep -x "${db}d" > /dev/null; then
        print_important "$db is running"
    fi
done
echo ""

# Application Directories
print_section "APPLICATION DIRECTORIES"
echo "Common application locations:"
for dir in /var/www /opt /home/*/apps /srv /usr/local/bin; do
    if [ -d "$dir" ]; then
        echo "$dir:"
        ls -la "$dir" 2>/dev/null | head -5
        echo ""
    fi
done

# PM2 processes (common for Node.js apps)
if command -v pm2 &> /dev/null; then
    print_section "PM2 PROCESSES"
    pm2 list 2>/dev/null || echo "PM2 installed but no processes or permission issues"
    echo ""
fi

# Node.js apps
if command -v node &> /dev/null; then
    print_section "NODE.JS PROCESSES"
    pgrep -f "node" | while read pid; do
        echo "PID $pid: $(ps -p $pid -o args --no-headers 2>/dev/null)"
    done
    echo ""
fi

# User accounts
print_section "USER ACCOUNTS"
echo "Non-system users:"
awk -F: '$3 >= 1000 {print $1 " (UID: " $3 ")"}' /etc/passwd
echo ""

# Cron jobs
print_section "SCHEDULED TASKS"
echo "System cron jobs:"
ls -la /etc/cron.* 2>/dev/null | grep -v "total"
echo ""
echo "User cron jobs:"
for user in $(awk -F: '$3 >= 1000 {print $1}' /etc/passwd); do
    if crontab -u "$user" -l 2>/dev/null | grep -v "^#" | grep -v "^$"; then
        echo "User $user has cron jobs"
    fi
done
echo ""

# SSL Certificates
print_section "SSL CERTIFICATES"
if [ -d "/etc/letsencrypt/live" ]; then
    echo "Let's Encrypt certificates:"
    ls -la /etc/letsencrypt/live/ 2>/dev/null
fi
if [ -d "/etc/ssl/certs" ]; then
    echo "SSL certificates in /etc/ssl/certs:"
    ls -la /etc/ssl/certs/*.crt 2>/dev/null | head -5
fi
echo ""

# Important Configuration Files
print_section "IMPORTANT CONFIG FILES"
for config in /etc/nginx/nginx.conf /etc/apache2/apache2.conf /etc/mysql/my.cnf /etc/postgresql/*/main/postgresql.conf; do
    if [ -f "$config" ]; then
        echo "Found: $config"
    fi
done
echo ""

# Recommendations
print_section "🎯 RECOMMENDATIONS FOR RESIST DEPLOYMENT"
echo ""
print_safe "SAFE APPROACH:"
echo "1. Create a backup of your entire server before making changes"
echo "2. Document your family chat app and CRM configurations"
echo "3. Use Docker containers for Resist blockchain (isolated from existing apps)"
echo "4. Use different ports for Resist services to avoid conflicts"
echo "5. Test Resist deployment in a separate user directory first"
echo ""

print_important "PORTS TO CHECK:"
echo "Family Chat App - Check what ports it uses"
echo "Business CRM - Document its port usage"
echo "Resist needs: 26656 (P2P), 26657 (RPC), 1317 (API), 3000 (Mobile API)"
echo ""

print_important "NEXT STEPS:"
echo "1. Run this audit and save the output"
echo "2. Identify your family chat app and CRM locations"
echo "3. Create a backup strategy"
echo "4. Plan port allocation to avoid conflicts"
echo "5. Use Docker for clean separation"
echo ""

echo "📊 AUDIT COMPLETE"
echo "Review this output carefully before proceeding with Resist deployment."
echo "Save this report: $0 > server-audit-$(date +%Y%m%d).txt"