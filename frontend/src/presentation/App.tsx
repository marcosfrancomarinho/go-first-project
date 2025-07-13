import React from 'react';
import { BrowserRouter } from 'react-router';
import { RouterView } from './routers/RouterView';

export const App: React.FC = () => {
  return (
    <BrowserRouter>
      <div className='flex flex-col min-h-screen'>
        <main className='grow'>
          <RouterView />
        </main>
      </div>
    </BrowserRouter>
  );
};
