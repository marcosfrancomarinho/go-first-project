import React, { type ReactNode } from 'react';
import type { AppConfig } from '../../shared/container/Container';

export const AppContext = React.createContext<AppConfig | null>(null);

export const AppProvider: React.FC<{
  children: ReactNode;
  dependecies: AppConfig;
}> = ({ children, dependecies }) => {
  return <AppContext.Provider value={dependecies }>{children}</AppContext.Provider>;
};
