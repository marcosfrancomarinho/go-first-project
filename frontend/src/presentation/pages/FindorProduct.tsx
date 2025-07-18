import React from 'react';
import type { ResponseFindorProductDTO } from '../../application/dto/ResponseFindorProductDTO';
import { AppContext } from '../hooks/Global';
import { Fallback } from '../components/Fallback';

export const FindorProduct: React.FC = () => {
  const { findorProductUseCase } = React.useContext(AppContext)!;
  const [products, setProducts] = React.useState<ResponseFindorProductDTO[] | null>(null);
  const [error, setError] = React.useState<Error | null>(null);
  const [loading, setLoading] = React.useState<boolean>(false);

  React.useEffect(() => {
    const request = async () => {
      try {
        setLoading(true);
        const productsList = await findorProductUseCase.findAll();
        setProducts(productsList);
      } catch (error) {
        setError(error as Error);
      } finally {
        setLoading(false);
      }
    };
    request();
  }, []);

  if (loading) return <Fallback loading={loading} />;
  if (error) return <Fallback error={error} />;
  return (
    <div className='max-w-7xl mx-auto px-4 py-10'>
      <h2 className='text-3xl font-bold text-blue-700 mb-6 text-center'>Produtos Cadastrados</h2>

      <div className='grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-6'>
        {products &&
          products.map((product) => (
            <div key={product.id} className='bg-white rounded-lg shadow-md p-6 border border-gray-100 hover:shadow-lg transition'>
              <h3 className='text-xl font-semibold text-gray-800 mb-2'>{product.name}</h3>
              <p className='text-gray-600'>Preço: R$ {product.price.toFixed(2)}</p>
              <p className='text-gray-600'>Quantidade: {product.quantity}</p>
              <p className='text-gray-800 font-medium mt-2'>Total: R$ {product.total.toFixed(2)}</p>
            </div>
          ))}
      </div>
    </div>
  );
};
