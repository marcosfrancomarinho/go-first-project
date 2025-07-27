import React from 'react';
import { AppContext } from '../context/Global';
import { useNavigate } from 'react-router';

export function useMenu() {
  const { authUserUseCase } = React.useContext(AppContext)!;
  const [name, setName] = React.useState('');
  const [authenticated, setAuthenticated] = React.useState(true);
  const navigate = useNavigate();

  const loadMenu = React.useCallback(() => {
    const isAuth = authUserUseCase.isAuthenticate();
    setAuthenticated(isAuth);
    if (isAuth) {
      const userName = authUserUseCase.getNameUser();
      setName(userName);
    }
  }, [authUserUseCase]);
  
  const logout = React.useCallback(() => {
    authUserUseCase.logoutUser();
    navigate('/');
  }, [authUserUseCase]);

  return { name, authenticated, loadMenu, logout };
}
