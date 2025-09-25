import React, { useState, useEffect, useCallback } from 'react';
import {
  View,
  FlatList,
  RefreshControl,
  StyleSheet,
  Alert,
} from 'react-native';
import {
  Card,
  Title,
  Paragraph,
  Button,
  Chip,
  ActivityIndicator,
  FAB,
} from 'react-native-paper';

import { useNode } from '../context/NodeContext';
import { Post } from '../services/ResistLiteNode';
import CreatePostModal from '../components/CreatePostModal';

const FeedScreen: React.FC = () => {
  const { liteNode } = useNode();
  const [posts, setPosts] = useState<Post[]>([]);
  const [loading, setLoading] = useState(false);
  const [refreshing, setRefreshing] = useState(false);
  const [hasMore, setHasMore] = useState(true);
  const [createPostVisible, setCreatePostVisible] = useState(false);

  const loadFeed = useCallback(async (refresh = false) => {
    if (!liteNode || loading) return;

    try {
      setLoading(true);
      const offset = refresh ? 0 : posts.length;

      const result = await liteNode.getFeed({
        limit: 20,
        offset,
        content_types: 'text,image',
      });

      if (refresh) {
        setPosts(result.posts);
      } else {
        setPosts(prev => [...prev, ...result.posts]);
      }

      setHasMore(result.has_more);
    } catch (error) {
      console.error('Failed to load feed:', error);
      Alert.alert('Error', 'Failed to load feed. Please try again.');
    } finally {
      setLoading(false);
      setRefreshing(false);
    }
  }, [liteNode, posts.length, loading]);

  const onRefresh = useCallback(() => {
    setRefreshing(true);
    loadFeed(true);
  }, [loadFeed]);

  const loadMore = useCallback(() => {
    if (hasMore && !loading) {
      loadFeed();
    }
  }, [hasMore, loading, loadFeed]);

  useEffect(() => {
    loadFeed(true);
  }, []);

  const formatTimeAgo = (timestamp: number): string => {
    const now = Math.floor(Date.now() / 1000);
    const diff = now - timestamp;

    if (diff < 60) return 'Just now';
    if (diff < 3600) return `${Math.floor(diff / 60)}m ago`;
    if (diff < 86400) return `${Math.floor(diff / 3600)}h ago`;
    return `${Math.floor(diff / 86400)}d ago`;
  };

  const getIntentColor = (intent: string): string => {
    switch (intent) {
      case 'educate': return '#4CAF50';
      case 'discuss': return '#2196F3';
      case 'share': return '#FF9800';
      case 'question': return '#9C27B0';
      default: return '#757575';
    }
  };

  const renderPost = ({ item }: { item: Post }) => (
    <Card style={styles.postCard}>
      <Card.Content>
        <View style={styles.postHeader}>
          <Title numberOfLines={2}>{item.title}</Title>
          <Chip
            mode="outlined"
            textStyle={{ color: getIntentColor(item.intent) }}
            style={{ borderColor: getIntentColor(item.intent) }}
          >
            {item.intent}
          </Chip>
        </View>

        <Paragraph style={styles.author}>
          by {item.author.slice(0, 12)}... • {formatTimeAgo(item.created_at)}
        </Paragraph>

        <Paragraph numberOfLines={4}>{item.content}</Paragraph>

        {item.sources.length > 0 && (
          <View style={styles.sourcesContainer}>
            <Paragraph style={styles.sourcesLabel}>Sources:</Paragraph>
            {item.sources.slice(0, 2).map((source, index) => (
              <Chip key={index} mode="outlined" compact>
                Source {index + 1}
              </Chip>
            ))}
          </View>
        )}
      </Card.Content>

      <Card.Actions>
        <Button icon="thumb-up" mode="outlined" compact>
          {item.upvotes}
        </Button>
        <Button icon="thumb-down" mode="outlined" compact>
          {item.downvotes}
        </Button>
        <Button icon="share" mode="outlined" compact>
          Share
        </Button>
      </Card.Actions>
    </Card>
  );

  const renderFooter = () => {
    if (!loading || refreshing) return null;
    return (
      <View style={styles.loadingFooter}>
        <ActivityIndicator size="small" />
      </View>
    );
  };

  const handleCreatePost = async (postData: any) => {
    try {
      await liteNode?.createPost(postData);
      setCreatePostVisible(false);
      onRefresh(); // Refresh feed to show new post
      Alert.alert('Success', 'Post created successfully!');
    } catch (error) {
      console.error('Failed to create post:', error);
      Alert.alert('Error', 'Failed to create post. Please try again.');
    }
  };

  return (
    <View style={styles.container}>
      <FlatList
        data={posts}
        renderItem={renderPost}
        keyExtractor={item => item.id}
        refreshControl={
          <RefreshControl refreshing={refreshing} onRefresh={onRefresh} />
        }
        onEndReached={loadMore}
        onEndReachedThreshold={0.5}
        ListFooterComponent={renderFooter}
        showsVerticalScrollIndicator={false}
      />

      <FAB
        style={styles.fab}
        icon="plus"
        onPress={() => setCreatePostVisible(true)}
      />

      <CreatePostModal
        visible={createPostVisible}
        onDismiss={() => setCreatePostVisible(false)}
        onSubmit={handleCreatePost}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  postCard: {
    margin: 8,
    elevation: 2,
  },
  postHeader: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'flex-start',
    marginBottom: 8,
  },
  author: {
    fontSize: 12,
    color: '#666',
    marginBottom: 8,
  },
  sourcesContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    flexWrap: 'wrap',
    marginTop: 8,
    gap: 4,
  },
  sourcesLabel: {
    fontSize: 12,
    fontWeight: 'bold',
    marginRight: 8,
  },
  loadingFooter: {
    padding: 16,
    alignItems: 'center',
  },
  fab: {
    position: 'absolute',
    margin: 16,
    right: 0,
    bottom: 0,
    backgroundColor: '#2196F3',
  },
});

export default FeedScreen;