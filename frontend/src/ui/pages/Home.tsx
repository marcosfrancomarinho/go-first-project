import React from 'react';
import { Link } from 'react-router-dom';
export const Home: React.FC = () => {
  return (
    <div className='max-w-5xl mx-auto mt-10 px-4 text-center'>
      <h2 className='text-4xl font-bold text-blue-700 mb-4'>Bem-vindo à Tools Ltda</h2>
      <p className='text-gray-700 mb-8 text-lg'>
        Aqui você encontra as melhores ferramentas e soluções para o seu negócio. Faça login, cadastre-se ou gerencie seus
        produtos com facilidade.
      </p>

      <div className='flex flex-col md:flex-row justify-center gap-4'>
        <Link to='/login' className='bg-blue-600 text-white px-6 py-3 rounded-lg hover:bg-blue-700 transition'>
          Fazer Login
        </Link>
        <Link
          to='/sign'
          className='bg-white border border-blue-600 text-blue-600 px-6 py-3 rounded-lg hover:bg-blue-50 transition'
        >
          Criar Conta
        </Link>
      </div>

      <div className='mt-12'>
        <h3 className='text-2xl font-semibold mb-2'>Nossos Serviços</h3>
        <ul className='text-left max-w-md mx-auto space-y-2 text-gray-700'>
          <li>✅ Cadastro de produtos</li>
          <li>✅ Consulta de estoque</li>
          <li>✅ Sistema prático e rápido</li>
          <li>✅ Interface moderna e responsiva</li>
        </ul>
      </div>
    </div>
  );
};
