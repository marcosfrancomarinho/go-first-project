import React from 'react';
import { Fallback } from '../components/Fallback';
import { useFindorProduct } from '../hooks/useFindorProduct';
import { useDeleterProduct } from '../hooks/useDeleterProduct';

export const FindorProduct: React.FC = () => {
  const { error, loading, products, findorProduct, setError } = useFindorProduct();
  const { deleterProduct } = useDeleterProduct();
  React.useEffect(() => {
    findorProduct();
  }, [findorProduct]);

  const deleteProduct = async (id: string, index: number) => {
    const response = confirm('deseja excuir');
    if (response) await deleterProduct({ id });
    products?.splice(index, 1);
    if (products?.length === 0) setError(new Error('nenhum produto encontrado'));
  };

  if (loading) return <Fallback loading={loading} />;
  if (error) return <Fallback error={error} />;
  return (
    <div className='max-w-7xl mx-auto px-4 py-10'>
      <h2 className='text-3xl font-bold text-blue-700 mb-6 text-center'>Produtos Cadastrados</h2>

      <div className='grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6'>
        {products &&
          products.map((product, index) => (
            <div
              key={product.id}
              onClick={() => deleteProduct(product.id, index)}
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
    </div>
  );
};
