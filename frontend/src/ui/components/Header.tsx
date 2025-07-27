import React from 'react';
import { Menu } from './Menu';

export const Header: React.FC = () => {
  return (
    <header className="bg-blue-600 text-white shadow-md">
      <div className="container mx-auto flex items-center justify-between px-4 py-3">
        <h1 className="text-2xl font-bold">Tools Ltda</h1>
        <Menu />
      </div>
    </header>
  );
};
