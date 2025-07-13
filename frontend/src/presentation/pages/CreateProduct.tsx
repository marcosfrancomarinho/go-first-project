import React from 'react';
import { AppContext } from '../hooks/Global';
import { Spinner } from '../components/Spinner';
import { AlertError } from '../components/AlertError';
import { AlertSuccess } from '../components/AlertSucess';

type PayloadCreatorProduct = { name: string; price: number; quantity: number };

export const CreateProduct: React.FC = () => {
  const [error, setError] = React.useState<Error | null>(null);
  const [loadding, setLoading] = React.useState<boolean>(false);
  const { creatorProductUseCase } = React.useContext(AppContext)!;
  const [success, setSuccess] = React.useState<string>('');
  const form = React.useRef<HTMLFormElement | null>(null);
  const handleSubmit = async (e: any): Promise<void> => {
    try {
      e.preventDefault();

      setLoading(true);
      setError(null);
      setSuccess('');
      const payload = Object.fromEntries(new FormData(e.target)) as unknown as PayloadCreatorProduct;
      const { message } = await creatorProductUseCase.create({
        name: payload.name,
        price: Number(payload.price),
        quantity: Number(payload.quantity),
      });
      setSuccess(message);
      form.current?.reset();
      const input = form.current?.elements.item(0) as HTMLInputElement;
      input.focus();
    } catch (error) {
      setError(error as Error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className='flex relative justify-center items-center min-h-screen bg-gray-50 px-4'>
      {success && <AlertSuccess message={success} />}
      {error && <AlertError message={error.message} />}
      <div className='bg-white p-8 rounded-xl shadow-md w-full max-w-xl'>
        <h2 className='text-2xl font-bold text-center text-blue-600 mb-6'>Criar Produto - Tools Ltda</h2>

        <form ref={form} onSubmit={(e) => handleSubmit(e)} className='space-y-4'>
          <div>
            <label htmlFor='name' className='block text-sm font-medium text-gray-700'>
              Nome do Produto
            </label>
            <input
              name='name'
              id='name'
              type='text'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <div>
            <label htmlFor='price' className='block text-sm font-medium text-gray-700'>
              Preço
            </label>
            <input
              name='price'
              id='price'
              type='number'
              step='0.01'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <div>
            <label htmlFor='quantity' className='block text-sm font-medium text-gray-700'>
              Quantidade
            </label>
            <input
              name='quantity'
              id='quantity'
              type='number'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <button
            type='submit'
            className='w-full bg-blue-600 flex justify-center items-center gap-2 text-white py-2 rounded-md hover:bg-blue-700 transition'
          >
            {loadding && <Spinner />}
            Cadastrar Produto
          </button>
        </form>
      </div>
    </div>
  );
};
