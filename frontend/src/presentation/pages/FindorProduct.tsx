import React from 'react';
import { Fallback } from '../components/Fallback';
import { useFindorProduct } from '../hooks/useFindorProduct';
import { useDeleterProduct } from '../hooks/useDeleterProduct';
import { ConfirmDialog } from '../components/ConfirmDialog';
import { useConfirmDeleteProduct } from '../hooks/useConfirmDeleteProduct';
import { AlertSuccess } from '../components/AlertSucess';
import { AlertError } from '../components/AlertError';
import { Pagination } from '../components/Pagination';

export const FindorProduct: React.FC = () => {
  const { error, loading, products, findorProduct, total, QUANTITY_OF_PRODUCT_PER_PAGE } = useFindorProduct();
  const { confirmDeleterProduct, messageSuccess, messageError, setMessageSuccess } = useDeleterProduct();
  const { choiseProductToRemove, id, index, showBox, setShowBox } = useConfirmDeleteProduct();
  const [page, setPage] = React.useState<number>(1);
  
  React.useEffect(() => {
    findorProduct({ page });
    setMessageSuccess('');
  }, [page]);

  const removeProduct = () => {
    confirmDeleterProduct({ id });
    setShowBox(false);
    products?.splice(index, 1);
  };

  if (loading) return <Fallback loading={loading} />;
  if (error) return <Fallback error={error} />;
  return (
    <div className='max-w-7xl mx-auto px-4 py-10 relative flex flex-col'>
      {messageError && <AlertError message={messageError.message} />}
      {!messageError && messageSuccess && <AlertSuccess message={messageSuccess} />}
      <h2 className='text-3xl font-bold text-blue-700 mb-6 text-center'>Produtos Cadastrados</h2>
      {showBox && <ConfirmDialog onConfirm={removeProduct} message='Deseja Excluir' onCancel={() => setShowBox(false)} />}
      <div className='grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6'>
        {products &&
          products.map((product, index) => (
            <div
              key={product.id}
              onClick={() => choiseProductToRemove(product.id, index)}
              className='bg-white rounded-lg shadow-md p-6 border cursor-pointer border-gray-100 hover:shadow-lg transition'
            >
              <h3 className='text-xl font-semibold text-gray-800 mb-2'>{product.name}</h3>
              <p className='text-gray-600'>
                Preço: {product.price.toLocaleString('pt-BR', { currency: 'BRL', style: 'currency' })}
              </p>
              <p className='text-gray-600'>Quantidade: {product.quantity}</p>
              <p className='text-gray-800 font-medium mt-2'>
                Total: {product.total.toLocaleString('pt-BR', { currency: 'BRL', style: 'currency' })}
              </p>
            </div>
          ))}
      </div>
      <Pagination
        limit={QUANTITY_OF_PRODUCT_PER_PAGE.current}
        onPageChange={(newPage) => setPage(newPage)}
        page={page}
        total={total}
      />
    </div>
  );
};
