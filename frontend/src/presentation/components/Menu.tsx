import React from 'react';
import { FiUser } from 'react-icons/fi';
import { Link } from 'react-router';
import { AppContext } from '../hooks/Global';
export const Menu: React.FC<{ authenticated: boolean }> = ({ authenticated }) => {
  const { name } = React.useContext(AppContext)!;

  return authenticated ? (
    <nav className='space-x-4'>
      <Link to='/home' className='hover:underline'>
        Home
      </Link>
      <Link to='/login' className='hover:underline'>
        Login
      </Link>
      <Link to='/sign' className='hover:underline'>
        Cadastrar
      </Link>
    </nav>
  ) : (
    <nav className='space-x-4 flex items-center'>
      <Link to='/auth/create-product' className='hover:underline'>
        Cadastrar Produtos
      </Link>
      <Link to='/auth/find-product' className='hover:underline'>
        Buscar Produtos
      </Link>
      <div className='flex items-center gap-2 text-white'>
        <FiUser />
        <span>
          Bem-vindo <span className='lowercase'>{name}</span>
        </span>
      </div>
    </nav>
  );
};
