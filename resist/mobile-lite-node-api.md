# Mobile Lite-Node API Specifications

## Overview

The Resist blockchain mobile lite-node API enables smartphones to participate as mini-nodes in the decentralized social network. This API provides efficient, battery-optimized access to blockchain functionality while maintaining security and decentralization principles.

## Core Features

### 1. Selective Synchronization
- Content filtering by topic, author, or engagement metrics
- Bandwidth-aware sync with WiFi vs cellular detection
- Offline-first architecture with local content caching

### 2. Signal Protocol Integration
- End-to-end encrypted messaging between nodes
- Forward secrecy with automatic key rotation
- Secure channel establishment for hub communication

### 3. Resource Contribution
- Mobile devices can offer storage/bandwidth when plugged in
- Dynamic resource allocation based on device capabilities
- Economic incentives for mobile node participation

## API Endpoints

### Authentication & Identity

#### POST /api/v1/auth/challenge
Request authentication challenge for mobile device.

```json
{
  "device_id": "mobile_device_12345",
  "public_key": "0x...",
  "device_info": {
    "platform": "iOS|Android",
    "version": "1.0.0",
    "capabilities": ["storage", "bandwidth", "compute"]
  }
}
```

Response:
```json
{
  "challenge": "random_challenge_string",
  "expires_at": 1640995200,
  "session_id": "session_abc123"
}
```

#### POST /api/v1/auth/verify
Verify challenge response and establish session.

```json
{
  "challenge": "random_challenge_string",
  "signature": "signed_challenge",
  "session_id": "session_abc123"
}
```

### Content & Posts

#### GET /api/v1/posts/feed
Get personalized content feed for mobile device.

Query parameters:
- `limit`: Number of posts (default: 20, max: 100)
- `offset`: Pagination offset
- `content_types`: Comma-separated list (text,image,video)
- `max_size`: Maximum total response size in bytes
- `since`: Unix timestamp for incremental sync
- `topics`: Comma-separated topic filters

Response:
```json
{
  "posts": [
    {
      "id": "post_123",
      "title": "Example Post",
      "content": "Post content...",
      "author": "cosmos1abc...",
      "created_at": 1640995200,
      "upvotes": 42,
      "downvotes": 3,
      "media_url": "ipfs://Qm...",
      "media_type": "image",
      "sources": ["source_1", "source_2"],
      "intent": "educate",
      "context_type": "fact-based",
      "size_bytes": 1024
    }
  ],
  "total_size": 20480,
  "has_more": true,
  "next_offset": 20,
  "sync_timestamp": 1640995300
}
```

#### POST /api/v1/posts/create
Create new post from mobile device.

```json
{
  "title": "Post Title",
  "content": "Post content",
  "media_data": "base64_encoded_media", // Optional
  "media_type": "image/jpeg",
  "sources": ["http://example.com/source"],
  "intent": "educate|discuss|share|question",
  "context_type": "fact-based|opinion|personal-experience|analysis",
  "offline_created": true,
  "local_timestamp": 1640995200
}
```

### Resource Management

#### GET /api/v1/node/status
Get current mobile node status and capabilities.

Response:
```json
{
  "node_id": "mobile_node_123",
  "status": "active|idle|charging|offline",
  "battery_level": 85,
  "charging": true,
  "network_type": "wifi|cellular|offline",
  "available_storage": 2048000000,
  "allocated_storage": 512000000,
  "bandwidth_limit": 1000000,
  "contributed_content": 15,
  "earnings_today": 250
}
```

#### POST /api/v1/node/resource-offer
Offer device resources to the network.

```json
{
  "storage_gb": 2,
  "bandwidth_mbps": 10,
  "duration_hours": 8,
  "price_per_hour": 10,
  "conditions": {
    "wifi_only": true,
    "charging_required": true,
    "idle_hours": ["22:00", "06:00"]
  }
}
```

### Synchronization

#### POST /api/v1/sync/selective
Initiate selective content synchronization.

```json
{
  "sync_profile": "minimal|standard|full",
  "content_filters": {
    "topics": ["technology", "science"],
    "authors": ["cosmos1abc..."],
    "engagement_threshold": 10,
    "max_age_days": 7
  },
  "network_constraints": {
    "wifi_only": true,
    "max_bandwidth": 1000000,
    "max_duration": 300
  }
}
```

### Signal Protocol Messaging

#### POST /api/v1/signal/establish-channel
Establish secure Signal protocol channel with another node.

```json
{
  "target_node": "node_456",
  "purpose": "content_sync|resource_coordination|general",
  "auto_rotate_keys": true
}
```

Response:
```json
{
  "channel_id": "channel_789",
  "established": true,
  "expires_at": 1641081600
}
```

#### POST /api/v1/signal/send
Send encrypted message via Signal protocol.

```json
{
  "channel_id": "channel_789",
  "message_type": "content_request|resource_offer|heartbeat",
  "payload": {
    "content_ids": ["post_123", "post_456"],
    "priority": "high|medium|low"
  }
}
```

## Mobile-Specific Optimizations

### Battery Optimization
- Background sync only when charging or high battery (>50%)
- Adaptive refresh rates based on user activity
- Intelligent content prefetching during WiFi availability

### Bandwidth Management
- Compression for all API responses
- Image/video quality adaptation based on network
- Incremental sync with merkle tree verification

### Storage Efficiency
- Content deduplication across cached posts
- Automatic cleanup of old/unused content
- Smart caching based on user engagement patterns

### Offline Capability
- Local SQLite database for cached content
- Offline post creation with background sync
- Conflict resolution for offline operations

## Security Features

### Device Authentication
- Hardware-backed key storage (iOS Keychain, Android Keystore)
- Biometric authentication for sensitive operations
- Device integrity verification

### Data Protection
- All API communications over TLS 1.3
- Signal protocol for peer-to-peer messaging
- Local database encryption at rest

### Privacy Preservation
- Minimal data collection and retention
- Optional tor/VPN integration for anonymity
- User-controlled analytics and telemetry

## Economic Model

### Earning Opportunities for Mobile Nodes
- **Storage Provider**: Earn tokens for storing content (1-5 tokens/GB/day)
- **Bandwidth Relay**: Earn tokens for relaying content (0.1 tokens/MB)
- **Content Curator**: Earn tokens for quality voting and moderation
- **Network Validator**: Light validation rewards for uptime

### Payment Integration
- Built-in wallet for receiving rewards
- Micro-payment channels for instant transactions
- Integration with mobile payment systems

## Technical Implementation Notes

### React Native Components
```javascript
// Example React Native integration
import { ResistLiteNode } from '@resist/lite-node';

const liteNode = new ResistLiteNode({
  nodeId: 'mobile_device_12345',
  apiEndpoint: 'https://api.resist.network',
  signalEndpoint: 'wss://signal.resist.network'
});

// Sync content in background
await liteNode.syncContent({
  profile: 'minimal',
  wifiOnly: true
});
```

### Performance Targets
- Initial sync: <30 seconds on WiFi
- Background sync: <5MB per hour
- Battery impact: <2% per day with normal usage
- Offline capability: 7 days of cached content

### Compatibility
- iOS 13.0+ (iPhone 6s and newer)
- Android API 21+ (Android 5.0+)
- 2GB RAM minimum, 4GB recommended
- 1GB free storage for content cache

This API specification enables mobile devices to participate meaningfully in the Resist decentralized social network while maintaining excellent user experience and device performance.