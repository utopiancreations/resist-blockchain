import AsyncStorage from '@react-native-async-storage/async-storage';
import NetInfo from '@react-native-community/netinfo';
import BackgroundJob from 'react-native-background-job';
import EncryptedStorage from 'react-native-encrypted-storage';

export interface LiteNodeConfig {
  nodeId: string;
  apiEndpoint: string;
  signalEndpoint: string;
  maxStorageGB?: number;
  maxBandwidthMbps?: number;
}

export interface Post {
  id: string;
  title: string;
  content: string;
  author: string;
  created_at: number;
  upvotes: number;
  downvotes: number;
  media_url?: string;
  sources: string[];
  intent: string;
  context_type: string;
}

export interface NodeStatus {
  node_id: string;
  status: 'active' | 'idle' | 'charging' | 'offline';
  battery_level: number;
  charging: boolean;
  network_type: 'wifi' | 'cellular' | 'offline';
  available_storage: number;
  allocated_storage: number;
  contributed_content: number;
  earnings_today: number;
}

export class ResistLiteNode {
  private config: LiteNodeConfig;
  private authToken: string | null = null;
  private isInitialized = false;
  private backgroundSyncEnabled = false;

  constructor(config: LiteNodeConfig) {
    this.config = {
      maxStorageGB: 1,
      maxBandwidthMbps: 10,
      ...config,
    };
  }

  async initialize(): Promise<void> {
    if (this.isInitialized) return;

    try {
      // Load stored auth token
      this.authToken = await EncryptedStorage.getItem('auth_token');

      // Initialize local database
      await this.initializeLocalStorage();

      // Set up background sync
      await this.setupBackgroundSync();

      // Start periodic tasks
      this.startPeriodicTasks();

      this.isInitialized = true;
      console.log('Resist Lite Node initialized successfully');
    } catch (error) {
      console.error('Failed to initialize Resist Lite Node:', error);
      throw error;
    }
  }

  // Authentication Methods
  async requestChallenge(deviceId: string, publicKey: string): Promise<any> {
    const response = await fetch(`${this.config.apiEndpoint}/api/v1/auth/challenge`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        device_id: deviceId,
        public_key: publicKey,
        device_info: {
          platform: 'mobile',
          version: '1.0.0',
          capabilities: ['storage', 'bandwidth']
        }
      })
    });
    return response.json();
  }

  async verifyChallenge(challenge: string, signature: string, sessionId: string): Promise<boolean> {
    try {
      const response = await fetch(`${this.config.apiEndpoint}/api/v1/auth/verify`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          challenge,
          signature,
          session_id: sessionId
        })
      });

      const result = await response.json();

      if (result.auth_token) {
        this.authToken = result.auth_token;
        await EncryptedStorage.setItem('auth_token', result.auth_token);
        return true;
      }

      return false;
    } catch (error) {
      console.error('Challenge verification failed:', error);
      return false;
    }
  }

  // Content Methods
  async getFeed(options: {
    limit?: number;
    offset?: number;
    content_types?: string;
    topics?: string;
  } = {}): Promise<{ posts: Post[]; has_more: boolean }> {
    if (!this.authToken) throw new Error('Not authenticated');

    const params = new URLSearchParams({
      limit: (options.limit || 20).toString(),
      offset: (options.offset || 0).toString(),
      max_size: '1048576', // 1MB
      ...(options.content_types && { content_types: options.content_types }),
      ...(options.topics && { topics: options.topics }),
    });

    const response = await fetch(`${this.config.apiEndpoint}/api/v1/posts/feed?${params}`, {
      headers: { 'Authorization': `Bearer ${this.authToken}` }
    });

    const data = await response.json();

    // Cache posts locally
    await this.cachePostsLocally(data.posts);

    return {
      posts: data.posts,
      has_more: data.has_more
    };
  }

  async createPost(post: {
    title: string;
    content: string;
    sources: string[];
    intent: string;
    context_type: string;
  }): Promise<Post> {
    if (!this.authToken) throw new Error('Not authenticated');

    const response = await fetch(`${this.config.apiEndpoint}/api/v1/posts/create`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.authToken}`
      },
      body: JSON.stringify({
        ...post,
        offline_created: false,
        local_timestamp: Date.now()
      })
    });

    return response.json();
  }

  // Node Management Methods
  async getNodeStatus(): Promise<NodeStatus> {
    if (!this.authToken) throw new Error('Not authenticated');

    const response = await fetch(`${this.config.apiEndpoint}/api/v1/node/status`, {
      headers: { 'Authorization': `Bearer ${this.authToken}` }
    });

    return response.json();
  }

  async offerResources(offer: {
    storage_gb: number;
    bandwidth_mbps: number;
    duration_hours: number;
    price_per_hour: number;
  }): Promise<any> {
    if (!this.authToken) throw new Error('Not authenticated');

    const response = await fetch(`${this.config.apiEndpoint}/api/v1/node/resource-offer`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.authToken}`
      },
      body: JSON.stringify({
        ...offer,
        conditions: {
          wifi_only: true,
          charging_required: true,
          idle_hours: ['22:00', '06:00']
        }
      })
    });

    return response.json();
  }

  // Sync Methods
  async syncContent(options: {
    profile?: 'minimal' | 'standard' | 'full';
    wifi_only?: boolean;
  } = {}): Promise<void> {
    if (!this.authToken) throw new Error('Not authenticated');

    const netInfo = await NetInfo.fetch();

    // Check network constraints
    if (options.wifi_only && netInfo.type !== 'wifi') {
      console.log('Skipping sync: WiFi required but not available');
      return;
    }

    const syncRequest = {
      sync_profile: options.profile || 'minimal',
      content_filters: {
        topics: await this.getUserTopics(),
        max_age_days: 7,
        engagement_threshold: 5
      },
      network_constraints: {
        wifi_only: options.wifi_only || false,
        max_bandwidth: this.config.maxBandwidthMbps! * 1000000,
        max_duration: 300 // 5 minutes
      }
    };

    const response = await fetch(`${this.config.apiEndpoint}/api/v1/sync/selective`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.authToken}`
      },
      body: JSON.stringify(syncRequest)
    });

    const syncJob = await response.json();
    console.log('Sync initiated:', syncJob.job_id);
  }

  // Signal Protocol Methods
  async establishSecureChannel(targetNode: string, purpose: string): Promise<string> {
    if (!this.authToken) throw new Error('Not authenticated');

    const response = await fetch(`${this.config.apiEndpoint}/api/v1/signal/establish-channel`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.authToken}`
      },
      body: JSON.stringify({
        target_node: targetNode,
        purpose,
        auto_rotate_keys: true
      })
    });

    const result = await response.json();
    return result.channel_id;
  }

  async sendSecureMessage(channelId: string, messageType: string, payload: any): Promise<void> {
    if (!this.authToken) throw new Error('Not authenticated');

    await fetch(`${this.config.apiEndpoint}/api/v1/signal/send`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${this.authToken}`
      },
      body: JSON.stringify({
        channel_id: channelId,
        message_type: messageType,
        payload
      })
    });
  }

  // Private Methods
  private async initializeLocalStorage(): Promise<void> {
    // Initialize SQLite database for offline content
    const dbExists = await AsyncStorage.getItem('db_initialized');
    if (!dbExists) {
      // Create local content cache
      await AsyncStorage.setItem('cached_posts', JSON.stringify([]));
      await AsyncStorage.setItem('user_preferences', JSON.stringify({
        topics: ['technology', 'science'],
        sync_profile: 'minimal',
        wifi_only: true
      }));
      await AsyncStorage.setItem('db_initialized', 'true');
    }
  }

  private async setupBackgroundSync(): Promise<void> {
    if (!this.backgroundSyncEnabled) {
      // Set up background task for periodic sync
      BackgroundJob.start({
        jobKey: 'resist_sync',
        period: 60000 * 30, // 30 minutes
      });

      BackgroundJob.on('resist_sync', () => {
        this.performBackgroundSync();
      });

      this.backgroundSyncEnabled = true;
    }
  }

  private async performBackgroundSync(): Promise<void> {
    try {
      const netInfo = await NetInfo.fetch();

      // Only sync on WiFi to preserve mobile data
      if (netInfo.type === 'wifi' && netInfo.isConnected) {
        await this.syncContent({ profile: 'minimal', wifi_only: true });
      }
    } catch (error) {
      console.error('Background sync failed:', error);
    }
  }

  private startPeriodicTasks(): void {
    // Periodic node status updates
    setInterval(() => {
      this.updateNodeMetrics();
    }, 60000 * 5); // Every 5 minutes
  }

  private async updateNodeMetrics(): Promise<void> {
    try {
      // Update local node metrics
      const metrics = {
        last_ping: Date.now(),
        content_served: await this.getContentServed(),
        earnings: await this.calculateEarnings()
      };

      await AsyncStorage.setItem('node_metrics', JSON.stringify(metrics));
    } catch (error) {
      console.error('Failed to update node metrics:', error);
    }
  }

  private async cachePostsLocally(posts: Post[]): Promise<void> {
    try {
      const existingCacheStr = await AsyncStorage.getItem('cached_posts');
      const existingCache: Post[] = existingCacheStr ? JSON.parse(existingCacheStr) : [];

      // Merge new posts with existing cache (avoid duplicates)
      const postIds = new Set(existingCache.map(p => p.id));
      const newPosts = posts.filter(p => !postIds.has(p.id));

      const updatedCache = [...existingCache, ...newPosts];

      // Keep only recent posts (last 7 days)
      const weekAgo = Date.now() / 1000 - (7 * 24 * 60 * 60);
      const recentPosts = updatedCache.filter(p => p.created_at > weekAgo);

      await AsyncStorage.setItem('cached_posts', JSON.stringify(recentPosts));
    } catch (error) {
      console.error('Failed to cache posts locally:', error);
    }
  }

  private async getUserTopics(): Promise<string[]> {
    const prefsStr = await AsyncStorage.getItem('user_preferences');
    const prefs = prefsStr ? JSON.parse(prefsStr) : {};
    return prefs.topics || ['technology', 'science'];
  }

  private async getContentServed(): Promise<number> {
    const metricsStr = await AsyncStorage.getItem('node_metrics');
    const metrics = metricsStr ? JSON.parse(metricsStr) : {};
    return metrics.content_served || 0;
  }

  private async calculateEarnings(): Promise<number> {
    // Mock earnings calculation
    const contentServed = await this.getContentServed();
    return contentServed * 0.1; // 0.1 tokens per piece of content served
  }
}