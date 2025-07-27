import React from 'react';
import { useNavigate } from 'react-router';
import { AppContext } from '../context/Global';
import type { RequestSignUserDTO } from '../../application/dto/RequestSignUserDTO';
export type PayloadFormSign = {
  name: string;
  email: string;
  password: string;
};

export const useSign = () => {
  const [error, setError] = React.useState<Error | null>(null);
  const [loading, setLoading] = React.useState<boolean>(false);
  const { signUserUseCase } = React.useContext(AppContext)!;
  const navigate = useNavigate();

  const signUser = React.useCallback(
    async (datas: PayloadFormSign) => {
      try {
        setLoading(true);
        setError(null);
        const payload: RequestSignUserDTO = { email: datas.email, name: datas.name, password: datas.password, path: '/register' };
        await signUserUseCase.sign(payload);
        navigate('/login', { state: { datas } });
      } catch (error: any) {
        setError(error)
      } finally {
        setLoading(false);
      }
    },
    [signUserUseCase]
  );

  return { error, loading, signUser };
};
