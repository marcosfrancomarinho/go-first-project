import React, { type ReactNode } from 'react';
import type { SignUserUseCase } from '../../application/usecase/SignUserUseCase';
import type { LoginUserUseCase } from '../../application/usecase/LoginUserUseCase';
import type { FindorProductUseCase } from '../../application/usecase/FindorProductUseCase';
import type { CreatorProductUseCase } from '../../application/usecase/CreatorProductUseCase';

interface Dependencies {
  signUserUseCase: SignUserUseCase;
  loginUserUseCase: LoginUserUseCase;
  findorProductUseCase: FindorProductUseCase;
  creatorProductUseCase: CreatorProductUseCase;
}

export interface AppConfig extends Dependencies {
  name: string;
  setName: (name: string) => void;
}

export const AppContext = React.createContext<AppConfig | null>(null);

export const AppProvider: React.FC<{
  children: ReactNode;
  dependecies: Dependencies;
}> = ({ children, dependecies }) => {
  const [name, setName] = React.useState<string>('');
  return <AppContext.Provider value={{ ...dependecies, name, setName }}>{children}</AppContext.Provider>;
};
