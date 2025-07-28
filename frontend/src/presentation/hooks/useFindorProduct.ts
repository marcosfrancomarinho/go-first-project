import React from 'react';
import { useNavigate } from 'react-router';
import type { ResponseFindorProductDTO } from '../../application/dto/ResponseFindorProductDTO';
import { AppContext } from '../context/Global';

export const useFindorProduct = () => {
  const { findorProductUseCase, authUserUseCase } = React.useContext(AppContext)!;
  const [products, setProducts] = React.useState<ResponseFindorProductDTO[] | null>(null);
  const [error, setError] = React.useState<Error | null>(null);
  const [loading, setLoading] = React.useState<boolean>(false);
  const navigate = useNavigate();

  const findorProduct = React.useCallback(async () => {
    try {
      setLoading(true);
      setError(null);
      const result = await findorProductUseCase.findAll();
      setProducts(result);
    } catch (error: any) {
      !authUserUseCase.isAuthenticate() && navigate('/');
      setError(error);
    } finally {
      setLoading(false);
    }
  }, [findorProductUseCase, , authUserUseCase, navigate]);

  return { products, error, loading, findorProduct, setError };
};
