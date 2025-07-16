import React from 'react';
import { Link } from 'react-router-dom';

export const NotFound: React.FC = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 px-4 text-center">
      <h1 className="text-7xl font-extrabold text-red-600">404</h1>
      <p className="mt-4 text-2xl font-semibold text-gray-800">Página não encontrada</p>
      <p className="mt-2 text-gray-600">
        A página que você está tentando acessar não existe ou foi removida.
      </p>
      <Link
        to="/"
        className="mt-6 px-6 py-2 text-white bg-blue-600 hover:bg-blue-700 rounded-lg transition"
      >
        Voltar para a Home
      </Link>
    </div>
  );
};
