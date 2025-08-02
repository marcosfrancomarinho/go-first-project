import React from 'react';
import { BrowserRouter } from 'react-router';
import { RouterView } from './presentation/routers/RouterView';
import { AppProvider } from './presentation/context/Global';
import { Container } from './shared/container/Container';
const container = Container.getInstance();

export const App: React.FC = () => {
  return (
    <AppProvider dependecies={container.dependencies()}>
      <BrowserRouter>
        <div className='flex flex-col min-h-screen'>
          <div className='grow'>
            <RouterView />
          </div>
        </div>
      </BrowserRouter>
    </AppProvider>
  );
};
