import React, { type ReactNode } from 'react';
import { Navigate } from 'react-router';
import { AppContext } from '../hooks/Global';

export const PrivateRouter: React.FC<{ children: ReactNode }> = ({ children }) => {
  const { authUserUseCase } = React.useContext(AppContext)!;
  if (authUserUseCase.isAuthenticate()) return children;
  return <Navigate to='/' replace />;
};
