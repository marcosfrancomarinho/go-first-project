import React from 'react';
import { useNavigate } from 'react-router';
import type { RequestLoginUserDTO } from '../../application/dto/RequestLoginUserDTO';
import { AppContext } from '../context/Global';
export type PayloadLogin = {
  email: string;
  password: string;
};

export const useLogin = () => {
  const navigate = useNavigate();
  const { loginUserUseCase } = React.useContext(AppContext)!;
  const [error, setError] = React.useState<Error | null>(null);
  const [loading, setLoading] = React.useState<boolean>(false);

  const loginUser = React.useCallback(
    async (datas: PayloadLogin) => {
      try {
        setLoading(true);
        setError(null);
        const payload: RequestLoginUserDTO = {
          email: datas.email,
          password: datas.password,
          path: '/login',
          key: 'user',
        };
        await loginUserUseCase.login(payload);
        navigate('/auth');
      } catch (error: any) {
        setError(error);
      } finally {
        setLoading(false);
      }
    },
    [loginUserUseCase]
  );

  return { loading, error, loginUser };
};
