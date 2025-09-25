import React, { createContext, useContext } from 'react';
import { ResistLiteNode } from '../services/ResistLiteNode';

interface NodeContextType {
  liteNode: ResistLiteNode | null;
}

const NodeContext = createContext<NodeContextType | undefined>(undefined);

export const useNode = (): NodeContextType => {
  const context = useContext(NodeContext);
  if (!context) {
    throw new Error('useNode must be used within a NodeProvider');
  }
  return context;
};

export const NodeProvider: React.FC<{
  children: React.ReactNode;
  liteNode: ResistLiteNode | null;
}> = ({ children, liteNode }) => {
  return (
    <NodeContext.Provider value={{ liteNode }}>
      {children}
    </NodeContext.Provider>
  );
};