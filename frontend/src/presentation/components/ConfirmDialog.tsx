import React from 'react';

interface ConfirmDialogProps {
  message: string;
  onConfirm: () => void;
  onCancel: () => void;
}

export const ConfirmDialog: React.FC<ConfirmDialogProps> = ({ message, onConfirm, onCancel }) => {
  return (
    <div className='fixed inset-0 bg-black/50 flex items-center justify-center z-50'>
      <div className='bg-white rounded-lg shadow-lg p-6 w-full max-w-sm'>
        <h2 className='text-lg font-semibold mb-4'>{message}</h2>
        <div className='flex justify-end gap-2'>
          <button onClick={onCancel} className='px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300'>
            Cancelar
          </button>
          <button onClick={onConfirm} className='px-4 py-2 bg-red-600 text-white rounded hover:bg-red-700'>
            Excluir
          </button>
        </div>
      </div>
    </div>
  );
};
