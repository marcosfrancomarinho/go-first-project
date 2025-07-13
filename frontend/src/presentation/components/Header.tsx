import React from 'react';
import { Menu } from './Menu';

export const Header: React.FC<{ authenticated?: boolean }> = ({ authenticated }) => {
  return (
    <header className='bg-blue-600 text-white shadow-md'>
      <div className='max-w-7xl mx-auto px-4 py-4 flex items-center justify-between'>
        <h1 className='text-2xl font-bold'>Tools Ltda</h1>
        <Menu authenticated={!authenticated} />
      </div>
    </header>
  );
};
