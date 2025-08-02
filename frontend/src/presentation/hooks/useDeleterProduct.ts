import React from 'react';
import { AppContext } from '../context/Global';
import type { RequestDeleterProductDTO } from '../../application/dto/RequestDeleterProductDTO';
import { useNavigate } from 'react-router';

export type PayloadDeleterProduct = {
  id: string;
  userName: string;
};

export const useDeleterProduct = () => {
  const [messageSuccess, setMessageSuccess] = React.useState<string>('');
  const [messageError, setMessageError] = React.useState<Error | null>(null);
  const { deleterProductUseCase, authUserUseCase } = React.useContext(AppContext)!;
  const navigate = useNavigate();
  const confirmDeleterProduct = React.useCallback(
    async (datas: PayloadDeleterProduct) => {
      try {
        setMessageSuccess('');
        setMessageError(null);
        const payload: RequestDeleterProductDTO = { id: datas.id };
        const message: string = `product ${datas.userName} excluido com sucesso.`;
        await deleterProductUseCase.delete(payload);
        setMessageSuccess(message);
      } catch (error: any) {
        !authUserUseCase.isAuthenticate() && navigate('/');
        setMessageError(error);
      }
    },
    [deleterProductUseCase]
  );

  return { confirmDeleterProduct, messageSuccess, messageError, setMessageSuccess, setMessageError };
};
