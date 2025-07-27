import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import { App } from './App';
import './ui/styles/global.css';

function main() {
  const rootElement = document.querySelector('#root');
  createRoot(rootElement!).render(
    <StrictMode>
      <App/>
    </StrictMode>
  );
}

main();
