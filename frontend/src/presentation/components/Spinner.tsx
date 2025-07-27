import React from 'react';
export const Spinner: React.FC = () => {
  return (
    <div className='flex justify-center items-center'>
      <div className='w-7 h-7 border-4 border-blue-500 border-t-transparent rounded-full animate-spin' />
    </div>
  );
};
