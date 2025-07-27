import React from 'react';
import { NavItem } from './NavItem';
import { useMenu } from '../hooks/useMenu';

export const Baseboard: React.FC = () => {
  const { authenticated, loadMenu, logout } = useMenu();
  React.useEffect(() => {
    loadMenu();
  }, []);

  return (
    <nav className='flex items-center gap-6 text-sm' aria-label='Menu de navegação'>
      {authenticated ? (
        <>
          <div className='flex gap-4'>
            <div onClick={logout} className='hover:underline transition cursor-pointer duration-200 ease-in-out'>
              Sair da Conta
            </div>
            <NavItem to='/auth/create-product' label='Cadastrar Produtos' />
            <NavItem to='/auth/find-product' label='Buscar Produtos' />
          </div>
        </>
      ) : (
        <>
          <NavItem to='/' label='Home' />
          <NavItem to='/login' label='Login' />
          <NavItem to='/sign' label='Cadastrar' />
        </>
      )}
    </nav>
  );
};
