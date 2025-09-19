# Resist Blockchain API Documentation

## Overview
This document outlines the comprehensive API endpoints for the Resist blockchain, designed to support a transformative social media platform focused on radical acceptance, restorative justice, and continued education.

## Core Philosophy
- **Radical Acceptance**: Open dialogue without censorship, but with education
- **Restorative Justice**: Community-driven resolution of content issues
- **Continued Education**: ML-powered source analysis and fact-checking
- **Truth and Peace**: Promoting factual discussions with proper citations

## REST API Endpoints

### Identity Module (Authentication & Profiles)

#### User Authentication
- `POST /resist/identity/v1/request-challenge` - Request authentication challenge
- `POST /resist/identity/v1/verify-signature` - Verify signature and authenticate

#### User Profiles
- `GET /resist/identity/v1/user-profile` - List all user profiles
- `GET /resist/identity/v1/user-profile/{address}` - Get specific user profile
- `POST /resist/identity/v1/user-profile` - Create user profile
- `PUT /resist/identity/v1/user-profile/{address}` - Update user profile
- `DELETE /resist/identity/v1/user-profile/{address}` - Delete user profile

### Posts Module (Social Media Content)

#### Social Posts
- `GET /resist/posts/v1/social-post` - List all social posts
- `GET /resist/posts/v1/social-post/{id}` - Get specific social post
- `POST /resist/posts/v1/social-post` - Create social post with sources
- `PUT /resist/posts/v1/social-post/{id}` - Update social post
- `DELETE /resist/posts/v1/social-post/{id}` - Delete social post

#### Voting System
- `GET /resist/posts/v1/vote` - List all votes
- `GET /resist/posts/v1/vote/{id}` - Get specific vote
- `POST /resist/posts/v1/vote-post` - Vote on a post (upvote/downvote)

#### Source Citations
- `GET /resist/posts/v1/source` - List all sources
- `GET /resist/posts/v1/source/{id}` - Get specific source with analysis
- `POST /resist/posts/v1/source` - Submit source for analysis
- `PUT /resist/posts/v1/source/{id}` - Update source credibility

#### Content Tagging & Discovery
- `GET /resist/posts/v1/post-tag` - List all post tags
- `GET /resist/posts/v1/post-tag/{id}` - Get specific tag with related posts
- `POST /resist/posts/v1/post-tag` - Create/update post tags

### UserGroups Module (Community Governance)

#### User Groups/DAOs
- `GET /resist/usergroups/v1/user-group` - List all user groups
- `GET /resist/usergroups/v1/user-group/{id}` - Get specific user group
- `POST /resist/usergroups/v1/user-group` - Create user group
- `PUT /resist/usergroups/v1/user-group/{id}` - Update user group

#### Governance Proposals
- `GET /resist/usergroups/v1/governance-proposal` - List all proposals
- `GET /resist/usergroups/v1/governance-proposal/{id}` - Get specific proposal
- `POST /resist/usergroups/v1/governance-proposal` - Create governance proposal
- `PUT /resist/usergroups/v1/governance-proposal/{id}` - Update proposal votes

#### Content Moderation (Restorative Justice)
- `GET /resist/usergroups/v1/content-report` - List content reports
- `GET /resist/usergroups/v1/content-report/{id}` - Get specific report
- `POST /resist/usergroups/v1/content-report` - Report content for community review
- `PUT /resist/usergroups/v1/content-report/{id}` - Update report resolution

### Rewards Module (Node Incentives)

#### Node Registration & Rewards
- `GET /resist/rewards/v1/node` - List registered nodes
- `POST /resist/rewards/v1/register-node` - Register node for rewards
- `GET /resist/rewards/v1/rewards/{address}` - Check node rewards

## Advanced Features

### ML-Powered Source Analysis
The platform integrates with AI systems to analyze sources:

1. **Credibility Scoring**: Sources receive scores based on factual accuracy
2. **Content Analysis**: ML analyzes source content for bias and reliability
3. **Cross-Reference Detection**: Automatically links related sources
4. **Fact-Checking Integration**: Connects with external fact-checking APIs

### Content Moderation Philosophy

Instead of censorship, the platform uses:

1. **Community Reporting**: Users report content with evidence
2. **Educational Response**: Community provides educational context
3. **Restorative Resolution**: Focus on learning and growth
4. **Source Requirements**: Controversial topics require source citations

### Post Intent Classification
Posts are classified by intent:
- `educate`: Sharing factual information
- `discuss`: Opening dialogue on topics
- `share`: Personal experiences or opinions
- `question`: Seeking information or clarification

### Context Types
- `fact-based`: Verifiable information with sources
- `opinion`: Personal viewpoints and perspectives
- `personal-experience`: Individual stories and experiences
- `analysis`: Interpretation of data or events

## Lite Node Architecture

### Mobile/Desktop Client Features
- **Offline Capability**: Cache recent posts and user data
- **Selective Sync**: Choose which content types to synchronize
- **Source Verification**: Verify source authenticity offline
- **Local Voting**: Queue votes for later submission

### API Rate Limiting
- Authentication required for all POST/PUT operations
- Read operations: 1000 requests/hour per IP
- Write operations: 100 requests/hour per authenticated user
- Source analysis: 50 requests/hour per user

### WebSocket Endpoints
Real-time updates for:
- `/ws/posts` - New posts and vote updates
- `/ws/governance` - Proposal updates and voting results
- `/ws/moderation` - Content reports and resolutions

## Error Handling

All endpoints return standard HTTP status codes:
- `200` - Success
- `400` - Invalid request parameters
- `401` - Authentication required
- `403` - Insufficient permissions
- `404` - Resource not found
- `429` - Rate limit exceeded
- `500` - Internal server error

Error responses include:
```json
{
  "error": "Error description",
  "code": "ERROR_CODE",
  "details": "Additional context"
}
```

## Authentication Flow

1. **Request Challenge**: Client requests authentication challenge
2. **Sign Challenge**: User signs challenge with private key
3. **Verify Signature**: Server verifies signature and creates session
4. **Profile Creation**: Verified users get automatic profile creation
5. **Session Management**: JWT tokens for subsequent requests

## Data Models

### Social Post with Sources
```json
{
  "id": "post123",
  "title": "Climate Change Discussion",
  "content": "Analysis of recent climate data...",
  "author": "resist1xyz...",
  "sources": ["source456", "source789"],
  "intent": "educate",
  "context_type": "fact-based",
  "upvotes": 42,
  "downvotes": 3,
  "requires_moderation": false,
  "created_at": 1672531200
}
```

### Source with ML Analysis
```json
{
  "id": "source456",
  "url": "https://climate.nasa.gov/...",
  "title": "NASA Climate Data",
  "credibility_score": 95,
  "analysis_summary": "High-credibility government source with peer-reviewed data",
  "verified": true
}
```

### Content Report (Restorative Justice)
```json
{
  "id": "report789",
  "post_id": "post123",
  "reporter": "resist1abc...",
  "reason": "Requires additional context",
  "evidence": "Source appears outdated, newer data available",
  "status": "under_review",
  "community_response": "Community provided updated sources",
  "resolution": "Educational resources added to post"
}
```

This API enables building applications that foster thoughtful discourse, evidence-based discussions, and community-driven moderation while maintaining the principles of radical acceptance and continued education.