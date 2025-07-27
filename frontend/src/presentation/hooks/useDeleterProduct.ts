import React from 'react';
import { AppContext } from '../context/Global';
import type { RequestDeleterProductDTO } from '../../application/dto/RequestDeleterProductDTO';

export type PayloadDeleterProduct = {
  id: string;
};

export const useDeleterProduct = () => {
  const [messageSuccess, setMessageSuccess] = React.useState<string>('');
  const [messageError, setMessageError] = React.useState<string>('');
  const { deleterProductUseCase } = React.useContext(AppContext)!;

  const deleterProduct = React.useCallback(
    async (datas: PayloadDeleterProduct) => {
      try {
        const payload: RequestDeleterProductDTO = { id: datas.id };
        const { message } = await deleterProductUseCase.delete(payload);
        setMessageSuccess(message);
      } catch (error: any) {
        setMessageError(error);
      }
    },
    [deleterProductUseCase]
  );

  return { deleterProduct, messageSuccess, messageError };
};
