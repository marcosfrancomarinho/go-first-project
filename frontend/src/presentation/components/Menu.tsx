import React from 'react';
import { FaUser } from 'react-icons/fa';
import { AppContext } from '../hooks/Global';
import { NavItem } from './NavItem';

export const Menu: React.FC = () => {
  const { authUserUseCase } = React.useContext(AppContext)!;
  const isAuthenticated = authUserUseCase.isAuthenticate();

  return (
    <nav className='flex items-center gap-6 text-sm' aria-label='Menu de navegação'>
      {isAuthenticated ? (
        <>
          <div className='flex gap-4'>
            <NavItem to='/auth' label='Home' />
            <NavItem to='/auth/create-product' label='Cadastrar Produtos' />
            <NavItem to='/auth/find-product' label='Buscar Produtos' />
          </div>
          <div className='flex items-center gap-2 bg-blue-700 px-3 py-1 rounded'>
            <FaUser />
            <span>
              Bem-vindo, <span className='lowercase'>{authUserUseCase.getNameUser()}</span>
            </span>
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


