import React, { type ReactNode } from 'react';
import type { CreatorProductUseCase } from '../../application/usecase/CreatorProductUseCase';
import type { FindorProductUseCase } from '../../application/usecase/FindorProductUseCase';
import type { LoginUserUseCase } from '../../application/usecase/LoginUserUseCase';
import type { SignUserUseCase } from '../../application/usecase/SignUserUseCase';
import type { AuthUserUseCase } from '../../application/usecase/AuthUserUseCase';

interface AppConfig {
  signUserUseCase: SignUserUseCase;
  loginUserUseCase: LoginUserUseCase;
  findorProductUseCase: FindorProductUseCase;
  creatorProductUseCase: CreatorProductUseCase;
  authUserUseCase: AuthUserUseCase;
}

export const AppContext = React.createContext<AppConfig | null>(null);

export const AppProvider: React.FC<{
  children: ReactNode;
  dependecies: AppConfig;
}> = ({ children, dependecies }) => {
  return <AppContext.Provider value={dependecies }>{children}</AppContext.Provider>;
};
