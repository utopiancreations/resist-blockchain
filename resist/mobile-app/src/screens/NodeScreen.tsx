import React, { useState, useEffect } from 'react';
import {
  View,
  ScrollView,
  StyleSheet,
  Alert,
} from 'react-native';
import {
  Card,
  Title,
  Paragraph,
  Button,
  ProgressBar,
  Chip,
  Switch,
  Divider,
  List,
} from 'react-native-paper';

import { useNode } from '../context/NodeContext';
import { NodeStatus } from '../services/ResistLiteNode';

const NodeScreen: React.FC = () => {
  const { liteNode } = useNode();
  const [nodeStatus, setNodeStatus] = useState<NodeStatus | null>(null);
  const [resourceSharingEnabled, setResourceSharingEnabled] = useState(false);
  const [syncEnabled, setSyncEnabled] = useState(true);
  const [loading, setLoading] = useState(false);

  const loadNodeStatus = async () => {
    if (!liteNode) return;

    try {
      const status = await liteNode.getNodeStatus();
      setNodeStatus(status);
    } catch (error) {
      console.error('Failed to load node status:', error);
      Alert.alert('Error', 'Failed to load node status');
    }
  };

  useEffect(() => {
    loadNodeStatus();
    const interval = setInterval(loadNodeStatus, 30000); // Update every 30 seconds
    return () => clearInterval(interval);
  }, [liteNode]);

  const handleOfferResources = async () => {
    if (!liteNode) return;

    try {
      setLoading(true);
      await liteNode.offerResources({
        storage_gb: 1,
        bandwidth_mbps: 5,
        duration_hours: 8,
        price_per_hour: 10,
      });
      Alert.alert('Success', 'Resources offered to the network!');
      setResourceSharingEnabled(true);
      await loadNodeStatus();
    } catch (error) {
      console.error('Failed to offer resources:', error);
      Alert.alert('Error', 'Failed to offer resources');
    } finally {
      setLoading(false);
    }
  };

  const handleSync = async () => {
    if (!liteNode) return;

    try {
      setLoading(true);
      await liteNode.syncContent({ profile: 'standard', wifi_only: true });
      Alert.alert('Success', 'Content sync initiated!');
      await loadNodeStatus();
    } catch (error) {
      console.error('Failed to sync content:', error);
      Alert.alert('Error', 'Failed to sync content');
    } finally {
      setLoading(false);
    }
  };

  const formatBytes = (bytes: number): string => {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
  };

  const getStatusColor = (status: string): string => {
    switch (status) {
      case 'active': return '#4CAF50';
      case 'idle': return '#FF9800';
      case 'charging': return '#2196F3';
      case 'offline': return '#F44336';
      default: return '#757575';
    }
  };

  const getNetworkIcon = (networkType: string): string => {
    switch (networkType) {
      case 'wifi': return 'wifi';
      case 'cellular': return 'signal-cellular-alt';
      case 'offline': return 'signal-off';
      default: return 'help';
    }
  };

  if (!nodeStatus) {
    return (
      <View style={styles.container}>
        <Card style={styles.card}>
          <Card.Content>
            <Title>Loading Node Status...</Title>
          </Card.Content>
        </Card>
      </View>
    );
  }

  const storageUsedPercentage = (nodeStatus.allocated_storage / nodeStatus.available_storage) * 100;

  return (
    <ScrollView style={styles.container}>
      {/* Node Status Overview */}
      <Card style={styles.card}>
        <Card.Content>
          <Title>Mini-Node Status</Title>
          <View style={styles.statusRow}>
            <Chip
              icon="circle"
              mode="outlined"
              textStyle={{ color: getStatusColor(nodeStatus.status) }}
              style={{ borderColor: getStatusColor(nodeStatus.status) }}
            >
              {nodeStatus.status.toUpperCase()}
            </Chip>
            <Chip icon={getNetworkIcon(nodeStatus.network_type)} mode="outlined">
              {nodeStatus.network_type.toUpperCase()}
            </Chip>
          </View>
          <Paragraph style={styles.nodeId}>
            Node ID: {nodeStatus.node_id}
          </Paragraph>
        </Card.Content>
      </Card>

      {/* Device Status */}
      <Card style={styles.card}>
        <Card.Content>
          <Title>Device Status</Title>
          <View style={styles.statusItem}>
            <Paragraph>Battery Level</Paragraph>
            <View style={styles.progressContainer}>
              <ProgressBar
                progress={nodeStatus.battery_level / 100}
                color={nodeStatus.charging ? '#4CAF50' : '#FF9800'}
                style={styles.progressBar}
              />
              <Paragraph>{nodeStatus.battery_level}%</Paragraph>
            </View>
          </View>
          <View style={styles.statusItem}>
            <Paragraph>Charging</Paragraph>
            <Chip icon={nodeStatus.charging ? 'battery-charging' : 'battery'}>
              {nodeStatus.charging ? 'Yes' : 'No'}
            </Chip>
          </View>
        </Card.Content>
      </Card>

      {/* Storage Status */}
      <Card style={styles.card}>
        <Card.Content>
          <Title>Storage Status</Title>
          <View style={styles.statusItem}>
            <Paragraph>Available Storage</Paragraph>
            <Paragraph>{formatBytes(nodeStatus.available_storage)}</Paragraph>
          </View>
          <View style={styles.statusItem}>
            <Paragraph>Used by Network</Paragraph>
            <View style={styles.progressContainer}>
              <ProgressBar
                progress={storageUsedPercentage / 100}
                color="#2196F3"
                style={styles.progressBar}
              />
              <Paragraph>{formatBytes(nodeStatus.allocated_storage)}</Paragraph>
            </View>
          </View>
        </Card.Content>
      </Card>

      {/* Contribution Stats */}
      <Card style={styles.card}>
        <Card.Content>
          <Title>Network Contribution</Title>
          <List.Item
            title="Content Served"
            description={`${nodeStatus.contributed_content} items`}
            left={props => <List.Icon {...props} icon="share" />}
          />
          <List.Item
            title="Earnings Today"
            description={`${nodeStatus.earnings_today} tokens`}
            left={props => <List.Icon {...props} icon="coin" />}
          />
        </Card.Content>
      </Card>

      {/* Node Controls */}
      <Card style={styles.card}>
        <Card.Content>
          <Title>Node Controls</Title>

          <View style={styles.controlItem}>
            <View>
              <Paragraph style={styles.controlTitle}>Resource Sharing</Paragraph>
              <Paragraph style={styles.controlSubtitle}>
                Share storage and bandwidth with the network
              </Paragraph>
            </View>
            <Switch
              value={resourceSharingEnabled}
              onValueChange={setResourceSharingEnabled}
            />
          </View>

          <Divider style={styles.divider} />

          <View style={styles.controlItem}>
            <View>
              <Paragraph style={styles.controlTitle}>Auto Sync</Paragraph>
              <Paragraph style={styles.controlSubtitle}>
                Automatically sync content when on WiFi
              </Paragraph>
            </View>
            <Switch
              value={syncEnabled}
              onValueChange={setSyncEnabled}
            />
          </View>
        </Card.Content>

        <Card.Actions>
          <Button
            mode="outlined"
            onPress={handleOfferResources}
            disabled={loading || resourceSharingEnabled}
            loading={loading}
          >
            {resourceSharingEnabled ? 'Resources Shared' : 'Offer Resources'}
          </Button>
          <Button
            mode="contained"
            onPress={handleSync}
            disabled={loading}
            loading={loading}
          >
            Sync Now
          </Button>
        </Card.Actions>
      </Card>

      {/* Signal Protocol Status */}
      <Card style={styles.card}>
        <Card.Content>
          <Title>Secure Messaging</Title>
          <List.Item
            title="Active Channels"
            description="2 encrypted channels"
            left={props => <List.Icon {...props} icon="shield-lock" />}
          />
          <List.Item
            title="Messages Today"
            description="15 secure messages"
            left={props => <List.Icon {...props} icon="message-secure" />}
          />
        </Card.Content>
        <Card.Actions>
          <Button mode="outlined">
            Manage Channels
          </Button>
        </Card.Actions>
      </Card>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#f5f5f5',
  },
  card: {
    margin: 8,
    elevation: 2,
  },
  statusRow: {
    flexDirection: 'row',
    gap: 8,
    marginVertical: 8,
  },
  nodeId: {
    fontSize: 12,
    color: '#666',
    marginTop: 8,
  },
  statusItem: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginVertical: 4,
  },
  progressContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    flex: 0.5,
    gap: 8,
  },
  progressBar: {
    flex: 1,
    height: 6,
    borderRadius: 3,
  },
  controlItem: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    paddingVertical: 8,
  },
  controlTitle: {
    fontWeight: 'bold',
  },
  controlSubtitle: {
    fontSize: 12,
    color: '#666',
  },
  divider: {
    marginVertical: 8,
  },
});

export default NodeScreen;