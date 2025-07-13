import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { App } from './presentation/App';
import './presentation/styles/global.css';
import { AppProvider } from './presentation/hooks/Global';
import { Container } from './shared/container/Container';

function main() {
  const container = Container.getInstance();
  const rootElement = document.querySelector('#root');
  
  createRoot(rootElement!).render(
    <StrictMode>
      <AppProvider dependecies={container.dependencies()}>
        <App />
      </AppProvider>
    </StrictMode>
  );
}

main();
