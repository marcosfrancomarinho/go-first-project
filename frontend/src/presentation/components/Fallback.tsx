import React from 'react';
import { FiAlertCircle } from 'react-icons/fi';
import { Spinner } from './Spinner';

interface FallbackProps {
  error?: Error;
  loading?: boolean;
}

export const Fallback: React.FC<FallbackProps> = ({ error, loading }) => {
  if (loading) {
    return (
      <div className='flex flex-col items-center justify-center min-h-[200px] text-gray-600'>
        <Spinner />
        <p className='mt-2'>Carregando...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className='flex p-5 flex-col items-center justify-center text-red-600'>
        <FiAlertCircle className='text-4xl' />
        <p className='mt-2 font-medium'>Algo deu errado.</p>
        <code className='text-xs text-gray-500'>{error.message}</code>
      </div>
    );
  }

  return null;
};
