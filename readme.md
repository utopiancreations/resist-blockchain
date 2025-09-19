# Resist Blockchain

**Resist** is a revolutionary social media blockchain built on **Cosmos SDK** that implements **radical acceptance**, **restorative justice**, and **continued education** principles. Created with [Ignite CLI](https://ignite.com/cli), this blockchain enables a new paradigm for online discourse focused on the betterment of humankind, promoting peace and cooperation, and sharing truth.

## 🌟 Core Philosophy

### Radical Acceptance
- **No Traditional Censorship**: Content remains accessible with educational context
- **Open Dialogue**: All voices heard with proper source citation and community wisdom
- **Educational Focus**: Learning and understanding over punishment

### Restorative Justice
- **Community-Driven Moderation**: Collective responsibility for discourse quality
- **Evidence-Based Reporting**: Factual foundation for all moderation actions
- **Growth-Focused Resolutions**: Learning outcomes rather than punitive measures

### Continued Education
- **Source Citation Requirements**: All claims must be backed by credible sources
- **ML-Powered Analysis**: AI-assisted fact-checking and source credibility scoring
- **Community Wisdom Integration**: Collective intelligence guides discourse

## 🏗️ Architecture Overview

### Core Modules

#### 🔐 Identity Module
- **Challenge-Response Authentication**: No email/password required
- **Cryptographic Security**: ECDSA secp256k1 signature verification
- **Automatic Profile Creation**: Seamless onboarding process

#### 📝 Posts Module
- **Enhanced Social Posts**: Title, content, media with source citations
- **Intent Classification**: "educate", "discuss", "share", "question"
- **Context Types**: "fact-based", "opinion", "personal-experience", "analysis"
- **Voting System**: Community-driven upvote/downvote consensus
- **Source Integration**: ML-ready credibility scoring and analysis

#### 👥 UserGroups Module
- **Content Reporting**: Evidence-based community moderation
- **Governance Proposals**: Democratic decision-making for chain evolution
- **Restorative Responses**: Educational rather than punitive outcomes

#### 🏆 Rewards Module
- **Contribution Recognition**: Incentivizing quality discourse and fact-sharing
- **Community Building**: Rewarding constructive participation

### 🤖 ML Integration Points

#### Source Analysis System
- **Credibility Scoring**: 0-100 ML-powered reliability assessment
- **Content Analysis**: AI-powered source browsing and information digest
- **Verification Tracking**: Community and expert source validation

#### Content Clustering
- **Smart Tagging**: Automated content categorization
- **Similarity Scoring**: ML-powered content clustering for discussion centralization
- **Related Post Discovery**: Cross-reference system to prevent fragmentation

## 🚀 Getting Started

### Development Setup

```bash
# Start the blockchain in development mode
ignite chain serve

# Reset and start fresh (if needed)
ignite chain serve --verbose --reset-once
```

### API Development

The blockchain provides comprehensive **REST/gRPC APIs** for lite node development:

- **Identity**: Authentication and user profiles
- **Posts**: Social content creation, voting, and source citation
- **UserGroups**: Governance, reporting, and community moderation
- **Rewards**: Contribution tracking and incentivization

See `api-documentation.md` for complete endpoint specifications.

### Testing Status ✅

Comprehensive testing completed with full validation of:
- Authentication system functionality
- Voting and consensus mechanisms
- Content moderation workflow
- Governance proposal system
- Source citation and tagging
- API endpoint functionality

See `TEST-VALIDATION-REPORT.md` for detailed validation results.

## 🔗 Key Features

### For Users
- **Source-Required Posts**: Cite credible sources for claims
- **Intent-Driven Content**: Clear purpose classification for better understanding
- **Community Moderation**: Participate in restorative justice processes
- **Democratic Governance**: Vote on platform evolution proposals

### For Developers
- **Modular Architecture**: Independent, scalable blockchain modules
- **ML-Ready Integration**: Prepared endpoints for AI/ML system integration
- **Comprehensive APIs**: Full REST/gRPC support for any frontend
- **Real-Time Updates**: WebSocket architecture for live content

### For Communities
- **Truth-Focused Platform**: Evidence-based discourse with source verification
- **Peace-Promoting Design**: Restorative rather than punitive moderation
- **Educational Outcomes**: Learning and growth from conflicts
- **Collective Intelligence**: Community-driven wisdom and decision-making

## 📱 Frontend Development

### Lite Node Architecture
- **Offline Capability**: Selective content synchronization
- **Content Filtering**: User-defined topic and context preferences
- **Rate Limiting**: Fair usage controls for sustainable operation
- **Real-Time Sync**: WebSocket integration for live updates

### Recommended Stack
- **Vue.js Frontend**: Use `ignite scaffold vue` for quick setup
- **API Integration**: Comprehensive REST/gRPC endpoint documentation available
- **ML Integration**: Source analysis and content clustering APIs ready

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/resist@latest! | sudo bash
```
`username/resist` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/ignite/installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.com/invite/ignitecli)
