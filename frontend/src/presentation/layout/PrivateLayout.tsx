import React from 'react';
import { Header } from '../components/Header';
import { Footer } from '../components/Footer';
import { Outlet } from 'react-router';
export const PrivateLayout: React.FC = () => {
  return (
    <div className='flex flex-col min-h-screen'>
      <Header />
      <main className='grow'>
        <Outlet />
      </main>
      <Footer />
    </div>
  );
};
