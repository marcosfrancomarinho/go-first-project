import React from 'react';
type ParamsButton = { menuOpen: boolean; setMenuOpen: (value: boolean) => void };
export const Hamburguer: React.FC<ParamsButton> = ({ menuOpen, setMenuOpen }) => {
  return (
    <button
      className='md:hidden p-2 rounded bg-blue-700'
      aria-label='Alternar menu'
      aria-expanded={menuOpen}
      onClick={() => setMenuOpen(!menuOpen)}
    >
      {/* Ícone simples hamburger */}
      <svg
        className='w-6 h-6 text-white'
        fill='none'
        stroke='currentColor'
        viewBox='0 0 24 24'
        xmlns='http://www.w3.org/2000/svg'
      >
        {menuOpen ? (
          // Ícone X para fechar
          <path strokeLinecap='round' strokeLinejoin='round' strokeWidth={2} d='M6 18L18 6M6 6l12 12' />
        ) : (
          // Ícone hamburger
          <path strokeLinecap='round' strokeLinejoin='round' strokeWidth={2} d='M4 6h16M4 12h16M4 18h16' />
        )}
      </svg>
    </button>
  );
};
