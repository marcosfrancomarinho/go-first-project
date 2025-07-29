import React from 'react';
import { useNavigate } from 'react-router';
import type { RequestFindorProductDTO } from '../../application/dto/RequestFindorProductDTO';
import type { AllProducts } from '../../application/dto/ResponseFindorProductDTO';
import { AppContext } from '../context/Global';

type PayloadDeleterProduct = {
  page: number;
};

export const useFindorProduct = () => {
  const QUANTITY_OF_PRODUCT_PER_PAGE = React.useRef(10);
  const { findorProductUseCase, authUserUseCase } = React.useContext(AppContext)!;
  const [products, setProducts] = React.useState<AllProducts[] | null>(null);
  const [error, setError] = React.useState<Error | null>(null);
  const [loading, setLoading] = React.useState<boolean>(false);
  const [total, setTotal] = React.useState<number>(0);
  const navigate = useNavigate();
  const findorProduct = React.useCallback(
    async (data: PayloadDeleterProduct) => {
      try {
        setLoading(true);
        setError(null);
        const payload: RequestFindorProductDTO = {
          query: `?page=${data.page}&limit=${QUANTITY_OF_PRODUCT_PER_PAGE.current}`,
        };
        const response = await findorProductUseCase.findAll(payload);
        setProducts(response.all);
        setTotal(response.total);
      } catch (error: any) {
        !authUserUseCase.isAuthenticate() && navigate('/');
        setError(error);
      } finally {
        setLoading(false);
      }
    },
    [findorProductUseCase, , authUserUseCase, navigate]
  );

  return { products, error, loading, findorProduct, setError, setProducts, total, QUANTITY_OF_PRODUCT_PER_PAGE };
};
