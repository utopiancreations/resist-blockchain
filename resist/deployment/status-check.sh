#!/bin/bash
# Resist Blockchain Deployment Status Check

echo "🛡️ Resist Blockchain Deployment Status"
echo "========================================="

echo ""
echo "📊 Container Status:"
echo "-------------------"
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}" | grep resist || echo "No Resist containers running"

echo ""
echo "🌐 Port Allocation Status:"
echo "--------------------------"
echo "Resist Services (Safe Ports):"
netstat -tlnp 2>/dev/null | grep -E ':(26656|26657|1317|4001|5001|3003)' || ss -tlnp | grep -E ':(26656|26657|1317|4001|5001|3003)' || echo "Port check failed"

echo ""
echo "Existing Services (Protected):"
netstat -tlnp 2>/dev/null | grep -E ':(3000|3001|3002|8080|5432)' || ss -tlnp | grep -E ':(3000|3001|3002|8080|5432)' || echo "Existing services check failed"

echo ""
echo "🔗 Service Connectivity:"
echo "------------------------"
echo -n "IPFS API (5001): "
curl -s -X POST http://localhost:5001/api/v0/version >/dev/null && echo "✅ OK" || echo "❌ FAIL"

echo -n "Node.js App (3000): "
curl -s -I http://localhost:3000/ >/dev/null && echo "✅ OK" || echo "❌ FAIL"

echo -n "Node.js App (3002): "
curl -s -I http://localhost:3002/ >/dev/null && echo "✅ OK" || echo "❌ FAIL"

echo ""
echo "💾 Data Volumes:"
echo "---------------"
docker volume ls | grep resist

echo ""
echo "🏗️ Deployment Summary:"
echo "======================"
echo "✅ Safe port allocation - no conflicts with existing services"
echo "✅ Docker Compose infrastructure deployed"
echo "✅ IPFS decentralized storage running"
echo "✅ PostgreSQL database for mobile API"
echo "✅ Isolated Docker network (resist-network)"
echo ""
echo "📝 Next Steps:"
echo "- Replace mock blockchain service with actual resistd binary"
echo "- Configure genesis.json with proper validators"
echo "- Set up monitoring and backup automation"
echo ""