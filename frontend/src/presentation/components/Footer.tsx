import React from 'react';
import { Menu } from './Menu';
export const Footer: React.FC<{ authenticated?: boolean }> = ({ authenticated }) => {
  return (
    <footer className='bg-blue-600 text-white mt-10'>
      <div className='max-w-7xl mx-auto px-4 py-6 flex flex-col md:flex-row justify-between items-center text-sm'>
        <p className='mb-2 md:mb-0'>&copy; {new Date().getFullYear()} Tools Ltda. Todos os direitos reservados.</p>
        <div className='space-x-4'>
          <Menu authenticated={!authenticated} />
        </div>
      </div>
    </footer>
  );
};
