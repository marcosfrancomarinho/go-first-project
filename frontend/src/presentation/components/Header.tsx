import React from 'react';
import { Menu } from './Menu';
import { Hamburguer } from './Hamburguer';

export const Header: React.FC = () => {
  const [menuOpen, setMenuOpen] = React.useState(true);
  return (
    <header className='bg-blue-600 text-white shadow-md'>
      <div className='container mx-auto flex flex-col md:flex-row items-start md:items-center justify-between px-4 py-3'>
        <div className='flex w-full justify-between'>
          <h1 className='text-2xl font-bold'>Tools Ltda</h1>
          <Hamburguer menuOpen={menuOpen} setMenuOpen={setMenuOpen} />
        </div>
        <Menu menuOpen={menuOpen} />
      </div>
    </header>
  );
};
