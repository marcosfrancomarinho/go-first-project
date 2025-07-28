import React, { useEffect } from 'react';
import { useMenu } from '../hooks/useMenu';
import { NavItem } from './NavItem';
import { FaUser } from 'react-icons/fa';
type ParamsMenu = { menuOpen: boolean };

export const Menu: React.FC<ParamsMenu> = ({ menuOpen }) => {
  const { authenticated, loadMenu, name } = useMenu();

  useEffect(() => {
    loadMenu();
  }, []);

  return (
    <nav className='relative' aria-label='Menu de navegação'>
      <div
        className={`
          ${menuOpen ? 'block' : 'hidden'} 
          sm:flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 py-4 text-sm
        `}
      >
        {authenticated ? (
          <>
            <div className='flex flex-col sm:flex-row gap-4 sm:items-center'>
              <NavItem to='/auth' label='Home' />
              <NavItem to='/auth/create-product' label='Cadastrar Produtos' />
              <NavItem to='/auth/find-product' label='Buscar Produtos' />
              <div className='flex items-center gap-2 bg-blue-700 px-3 py-1 rounded self-start md:self-auto'>
                <FaUser />
                <span>
                  Bem-vindo, <span className='lowercase'>{name}</span>
                </span>
              </div>
            </div>
          </>
        ) : (
          <div className='flex flex-col sm:flex-row gap-4'>
            <NavItem to='/' label='Home' />
            <NavItem to='/login' label='Login' />
            <NavItem to='/sign' label='Cadastrar' />
          </div>
        )}
      </div>
    </nav>
  );
};
