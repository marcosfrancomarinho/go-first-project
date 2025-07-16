import React, { type ReactNode } from 'react';
import { Navigate } from 'react-router';
import { AppContext } from '../hooks/Global';

export const PublicRouter: React.FC<{ children: ReactNode }> = ({ children }) => {
  const { authUserUseCase } = React.useContext(AppContext)!;
  if (authUserUseCase.isAuthenticate()) return <Navigate to='/auth' replace />;
  return children;
};
