import React from 'react';

export const AlertError: React.FC<{ message: string }> = ({ message }) => {
  const [visible, setVisible] = React.useState<boolean>(true);

  React.useEffect(() => {
    setTimeout(() => setVisible(false), 3000);
  }, []);

  if (visible)
    return (
      <div className='absolute top-5 min-w-1 items-center gap-5 justify-between px-2 py-3 flex text-sm text-red-800 bg-red-100 border border-red-400 rounded-lg'>
        <div className='space-x-1'>
          <strong className='font-semibold'>Erro:</strong>
          <span>{message}</span>
        </div>
        <button onClick={() => setVisible(false)} className=' text-red-800 hover:text-red-600 text-lg font-bold'>
          &times;
        </button>
      </div>
    );
};
