import React from 'react';

type ParamsPagination = {
  page: number;
  limit: number;
  total: number;
  onPageChange: (page: number) => void;
};

export const Pagination: React.FC<ParamsPagination> = ({ page, limit, total, onPageChange }) => {
  const totalPages = Math.ceil(total / limit);

  const handlePageChange = (newPage: number) => {
    if (newPage < 1 || newPage > totalPages) return;
    onPageChange(newPage);
  };

  return (
    <div
      className='
        flex justify-center items-center gap-1 mt-4
        overflow-x-auto
        px-2
        scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-gray-100
      '
      style={{ WebkitOverflowScrolling: 'touch' }} 
    >
      <button
        onClick={() => handlePageChange(page - 1)}
        disabled={page === 1}
        className='
          px-2 py-1
          bg-gray-200 rounded
          hover:bg-gray-300
          disabled:opacity-50
          text-sm
          flex-shrink-0
        '
      >
        Anterior
      </button>

      {Array.from({ length: totalPages }, (_, i) => (
        <button
          key={i + 1}
          onClick={() => handlePageChange(i + 1)}
          className={`
            px-2 py-1 rounded text-sm flex-shrink-0
            ${page === i + 1 ? 'bg-blue-600 text-white' : 'bg-gray-100 hover:bg-gray-300'}
          `}
        >
          {i + 1}
        </button>
      ))}

      <button
        onClick={() => handlePageChange(page + 1)}
        disabled={page === totalPages}
        className='
          px-2 py-1
          bg-gray-200 rounded
          hover:bg-gray-300
          disabled:opacity-50
          text-sm
          flex-shrink-0
        '
      >
        Próxima
      </button>
    </div>
  );
};
