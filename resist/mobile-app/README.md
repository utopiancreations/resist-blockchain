# Resist Mobile App - Mini-Node Prototype

## Overview

The Resist mobile app transforms smartphones into mini-nodes in the decentralized social network. This React Native prototype demonstrates how mobile devices can participate in content distribution, resource sharing, and secure messaging using Signal protocol.

## Key Features

### 🔐 Decentralized Authentication
- No email/password required
- Challenge-response cryptographic authentication
- Hardware-backed key storage
- Automatic session management

### 📱 Mini-Node Functionality
- **Resource Sharing**: Contribute storage and bandwidth to the network
- **Content Caching**: Store frequently accessed content for offline use
- **Secure Messaging**: Signal protocol for encrypted node-to-node communication
- **Economic Participation**: Earn tokens for contributing resources

### 🌐 Smart Synchronization
- **WiFi-First**: Prioritize WiFi for data-intensive operations
- **Battery Aware**: Adjust activity based on charging status
- **Selective Sync**: Only sync relevant content based on user interests
- **Offline Capable**: Cache content for 7 days of offline browsing

### 🔒 Privacy & Security
- **End-to-End Encryption**: All communications use Signal protocol
- **Local Data Encryption**: SQLite database encrypted at rest
- **Minimal Data Collection**: Only essential data for network participation
- **Optional Anonymity**: Built-in support for Tor/VPN routing

## Technical Architecture

### Core Services

#### ResistLiteNode Service
- Main service class managing all blockchain interactions
- Handles authentication, content sync, and resource management
- Background task coordination for minimal battery impact

#### Signal Protocol Integration
- Secure channel establishment between nodes
- Message encryption/decryption with forward secrecy
- Automatic key rotation for enhanced security

#### Content Distribution
- IPFS integration for decentralized content storage
- Intelligent caching based on user engagement
- Conflict resolution for offline-created content

### Mobile Optimizations

#### Battery Management
- Background sync only when charging (>50% battery)
- Adaptive refresh rates based on user activity
- Intelligent prefetching during favorable conditions

#### Network Efficiency
- Compression for all API communications
- Adaptive media quality based on connection type
- Delta sync to minimize bandwidth usage

#### Storage Optimization
- Content deduplication across cached posts
- Automatic cleanup of old/unused content
- Smart caching prioritizing engaged content

## Installation & Setup

### Prerequisites
- Node.js 18+
- React Native development environment
- iOS Simulator or Android Emulator

### Quick Start

```bash
# Install dependencies
npm install

# Start Metro bundler
npm start

# Run on iOS
npm run ios

# Run on Android
npm run android
```

### Development Configuration

Create `.env` file:
```
RESIST_API_ENDPOINT=https://api.resist.network
RESIST_SIGNAL_ENDPOINT=wss://signal.resist.network
DEBUG_MODE=true
```

## Architecture Components

### Screens
- **FeedScreen**: Personalized social content feed with offline support
- **NodeScreen**: Mini-node status, resource management, and earnings
- **SettingsScreen**: User preferences, sync options, and privacy controls
- **AuthScreen**: Decentralized authentication flow

### Services
- **ResistLiteNode**: Core blockchain interaction service
- **AuthenticationService**: Cryptographic authentication handling
- **ContentService**: Feed management and post creation
- **SyncService**: Background content synchronization
- **SignalService**: Secure peer-to-peer messaging

### Context Providers
- **AuthContext**: User authentication state management
- **NodeContext**: Mini-node service provider

## Economic Model

### Earning Opportunities
- **Storage Provider**: 1-5 tokens/GB/day for storing content
- **Bandwidth Relay**: 0.1 tokens/MB for relaying content
- **Content Curation**: Rewards for quality voting and moderation
- **Network Validation**: Light validation rewards for uptime

### Resource Contributions
- Configurable storage allocation (1-5GB typical)
- Bandwidth sharing during idle periods
- Content relay for nearby users
- Participation in light validation

## Privacy Features

### Data Minimization
- No personal information collection
- Pseudonymous user identifiers
- Optional metadata obfuscation

### Security Measures
- Hardware security module integration
- Biometric authentication for sensitive operations
- Secure enclave for key storage
- Network traffic analysis protection

## Performance Targets

- **Initial Sync**: <30 seconds on WiFi
- **Background Sync**: <5MB per hour
- **Battery Impact**: <2% per day with normal usage
- **Offline Duration**: 7 days of cached content
- **Memory Usage**: <100MB typical, 200MB maximum

## Development Roadmap

### Phase 1: Core Functionality ✅
- Basic authentication and feed display
- Mini-node status monitoring
- Resource sharing interface

### Phase 2: Advanced Features (In Progress)
- Signal protocol implementation
- IPFS content distribution
- Advanced sync algorithms
- Economic reward system

### Phase 3: Production Ready
- Performance optimization
- Security audits
- App store deployment
- Multi-language support

## Contributing

### Code Structure
```
src/
├── components/     # Reusable UI components
├── screens/        # Screen components
├── services/       # Business logic services
├── context/        # React context providers
├── utils/          # Utility functions
└── types/          # TypeScript type definitions
```

### Development Guidelines
- Use TypeScript for all new code
- Follow React Native best practices
- Implement proper error handling
- Add unit tests for critical functions
- Document API interfaces

## Testing

```bash
# Run unit tests
npm test

# Run E2E tests
npm run test:e2e

# Test on device
npm run build
```

## License

This project is licensed under the MIT License - promoting open-source development of decentralized social networks.

## Contact & Support

- **Issues**: Report bugs and feature requests on GitHub
- **Discussions**: Join community discussions
- **Security**: Report security issues privately

---

*Building the future of decentralized social media, one mobile device at a time.*