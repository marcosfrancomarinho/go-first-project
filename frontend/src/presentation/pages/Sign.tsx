import React from 'react';
import { Link } from 'react-router-dom';
import { AlertError } from '../components/AlertError';
import { Spinner } from '../components/Spinner';
import { useSign, type PayloadFormSign } from '../hooks/useSign';

export const Sign: React.FC = () => {
  const { error, loading, signUser } = useSign();

  const submitHandler = async (e: any): Promise<void> => {
    e.preventDefault();
    const datas = Object.fromEntries(new FormData(e.target)) as PayloadFormSign;
    await signUser(datas);
  };

  return (
    <div className='flex relative justify-center items-center sm:min-h-screen bg-gray-50 px-4'>
      {error && <AlertError message={error.message} />}
      <div className='bg-white mt-5 p-8 rounded-xl shadow-md w-full max-w-md'>
        <h2 className='text-2xl font-bold text-center text-blue-600 mb-6'>Cadastro - Tools Ltda</h2>
        <form onSubmit={(e) => submitHandler(e)} className='space-y-4'>
          <div>
            <label htmlFor='name' className='block text-sm font-medium text-gray-700'>
              Nome
            </label>
            <input
              name='name'
              id='name'
              type='text'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <div>
            <label htmlFor='email' className='block text-sm font-medium text-gray-700'>
              E-mail
            </label>
            <input
              name='email'
              id='email'
              type='email'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <div>
            <label htmlFor='password' className='block text-sm font-medium text-gray-700'>
              Senha
            </label>
            <input
              name='password'
              id='password'
              type='password'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <button
            type='submit'
            className='w-full flex justify-center items-center gap-3 bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 transition'
          >
            {loading && <Spinner />}
            Criar Conta
          </button>
        </form>

        <p className='text-center text-sm text-gray-600 mt-4'>
          Já tem uma conta?{' '}
          <Link to='/login' className='text-blue-600 hover:underline'>
            Faça login
          </Link>
        </p>
      </div>
    </div>
  );
};
