import React from 'react';
import type { RequestCreatorProductDTO } from '../../application/dto/RequestCreatorProductDTO';
import { AppContext } from '../context/Global';
export type PayloadCreatorProduct = {
  name: string;
  price: number;
  quantity: number;
};

export const useCreatorProduct = () => {
  const [error, setError] = React.useState<Error | null>(null);
  const [loadding, setLoading] = React.useState<boolean>(false);
  const [success, setSuccess] = React.useState<string>('');
  const { creatorProductUseCase } = React.useContext(AppContext)!;

  const createorProduct = React.useCallback(
    async (datas: PayloadCreatorProduct) => {
      try {
        setLoading(true);
        setError(null);
        setSuccess('');
        const payload: RequestCreatorProductDTO = {
          name: datas.name,
          price: datas.price,
          quantity: datas.quantity,
          key: 'user',
          path: '/product',
        };
        const { message } = await creatorProductUseCase.create(payload);
        setSuccess(message);
      } catch (error: any) {
        setError(error);
      } finally {
        setLoading(false);
      }
    },
    [creatorProductUseCase]
  );

  return { error, loadding, success, createorProduct };
};
