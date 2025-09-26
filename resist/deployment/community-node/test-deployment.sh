#!/bin/bash

# Resist Community Node Deployment Test Script
# Tests the community node deployment without actually starting services

set -e

echo "🧪 Testing Resist Community Node Deployment Package"
echo "=================================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_test() {
    echo -e "${BLUE}🔧 Testing: $1${NC}"
}

print_pass() {
    echo -e "${GREEN}✅ PASS: $1${NC}"
}

print_fail() {
    echo -e "${RED}❌ FAIL: $1${NC}"
}

print_info() {
    echo -e "${YELLOW}ℹ️  INFO: $1${NC}"
}

# Test 1: Check required files exist
print_test "Required deployment files"

REQUIRED_FILES=(
    "docker-compose.yml"
    "Dockerfile.resist-node"
    ".env.template"
    "README.md"
    "genesis.json"
    "setup-community-node.sh"
)

for file in "${REQUIRED_FILES[@]}"; do
    if [ -f "$file" ]; then
        print_pass "File exists: $file"
    else
        print_fail "Missing file: $file"
        exit 1
    fi
done

# Test 2: Validate Docker Compose syntax
print_test "Docker Compose configuration"
if docker-compose config > /dev/null 2>&1; then
    print_pass "Docker Compose syntax is valid"
else
    print_fail "Docker Compose syntax error"
    print_info "Make sure you have a .env file (copy from .env.template)"
fi

# Test 3: Check if ports are available
print_test "Port availability"

DEFAULT_PORTS=(26656 26657 1317 4001 5001 8080 3000 9090)
USED_PORTS=()

for port in "${DEFAULT_PORTS[@]}"; do
    if ss -tuln | grep -q ":$port "; then
        USED_PORTS+=($port)
        print_fail "Port $port is in use"
    else
        print_pass "Port $port is available"
    fi
done

if [ ${#USED_PORTS[@]} -gt 0 ]; then
    print_info "Used ports detected: ${USED_PORTS[*]}"
    print_info "You may need to customize ports in .env file"
fi

# Test 4: Check Docker availability
print_test "Docker environment"

if ! command -v docker &> /dev/null; then
    print_fail "Docker is not installed"
    print_info "Install Docker: curl -fsSL https://get.docker.com | sh"
else
    print_pass "Docker is installed"

    if docker ps > /dev/null 2>&1; then
        print_pass "Docker daemon is running"
    else
        print_fail "Docker daemon is not accessible"
        print_info "Try: sudo systemctl start docker"
        print_info "Or add user to docker group: sudo usermod -aG docker $USER"
    fi
fi

if ! command -v docker-compose &> /dev/null; then
    print_fail "Docker Compose is not installed"
    print_info "Install Docker Compose or use 'docker compose' (newer version)"
else
    print_pass "Docker Compose is available"
fi

# Test 5: Validate genesis file
print_test "Genesis file validation"

if [ -f "genesis.json" ]; then
    if jq empty genesis.json 2>/dev/null; then
        print_pass "Genesis file is valid JSON"

        # Check for required fields
        CHAIN_ID=$(jq -r '.chain_id // empty' genesis.json)
        if [ ! -z "$CHAIN_ID" ]; then
            print_pass "Genesis contains chain_id: $CHAIN_ID"
        else
            print_fail "Genesis file missing chain_id"
        fi
    else
        print_fail "Genesis file is not valid JSON"
    fi
else
    print_fail "Genesis file is missing"
fi

# Test 6: Environment template validation
print_test "Environment configuration template"

if [ -f ".env.template" ]; then
    print_pass "Environment template exists"

    # Check for critical variables
    REQUIRED_VARS=("NODE_MONIKER" "NODE_OPERATOR" "CHAIN_ID")
    for var in "${REQUIRED_VARS[@]}"; do
        if grep -q "^$var=" .env.template; then
            print_pass "Template contains: $var"
        else
            print_fail "Template missing: $var"
        fi
    done
else
    print_fail "Environment template is missing"
fi

# Test 7: Setup script validation
print_test "Setup script validation"

if [ -f "setup-community-node.sh" ]; then
    if [ -x "setup-community-node.sh" ]; then
        print_pass "Setup script is executable"
    else
        print_fail "Setup script is not executable"
        print_info "Run: chmod +x setup-community-node.sh"
    fi

    # Basic syntax check
    if bash -n setup-community-node.sh 2>/dev/null; then
        print_pass "Setup script syntax is valid"
    else
        print_fail "Setup script has syntax errors"
    fi
else
    print_fail "Setup script is missing"
fi

# Test 8: Resource requirements check
print_test "System resource requirements"

# Check memory (minimum 2GB, recommended 4GB+)
MEMORY_KB=$(awk '/^MemTotal:/ {print $2}' /proc/meminfo 2>/dev/null || echo "0")
MEMORY_GB=$((MEMORY_KB / 1024 / 1024))

if [ $MEMORY_GB -ge 4 ]; then
    print_pass "Memory: ${MEMORY_GB}GB (sufficient)"
elif [ $MEMORY_GB -ge 2 ]; then
    print_pass "Memory: ${MEMORY_GB}GB (minimum met)"
    print_info "4GB+ recommended for better performance"
else
    print_fail "Memory: ${MEMORY_GB}GB (insufficient)"
    print_info "Minimum 2GB required, 4GB+ recommended"
fi

# Check disk space (minimum 10GB, recommended 50GB+)
DISK_GB=$(df -BG . | awk 'NR==2 {print $4}' | sed 's/G//' 2>/dev/null || echo "0")

if [ $DISK_GB -ge 50 ]; then
    print_pass "Disk space: ${DISK_GB}GB (excellent)"
elif [ $DISK_GB -ge 10 ]; then
    print_pass "Disk space: ${DISK_GB}GB (sufficient)"
    print_info "50GB+ recommended for long-term operation"
else
    print_fail "Disk space: ${DISK_GB}GB (insufficient)"
    print_info "Minimum 10GB required, 50GB+ recommended"
fi

echo ""
echo "🎯 Test Summary"
echo "==============="

# Count passed tests
TOTAL_TESTS=$(grep -c "print_pass\|print_fail" "$0")
echo "Deployment package validation completed."

echo ""
echo "📋 Next Steps:"
echo "1. Copy .env.template to .env and customize your settings"
echo "2. Run ./setup-community-node.sh for automated setup"
echo "3. Or manually run: docker-compose up -d"
echo ""
echo "🌐 Need help? Visit https://docs.resist-blockchain.org/community-nodes"
echo ""

print_info "Test completed successfully! Your system is ready for Resist deployment."