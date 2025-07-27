import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { PrivateLayout } from '../layout/PrivateLayout';
import { PublicLayout } from '../layout/PublicLayout';
import { CreateProduct } from '../pages/CreateProduct';
import { FindorProduct } from '../pages/FindorProduct';
import { Home } from '../pages/Home';
import { Login } from '../pages/Login';
import { NotFound } from '../pages/NotFound';
import { Sign } from '../pages/Sign';
import { Welcome } from '../pages/Welcome';
import { PrivateRouter } from './PrivateRouter';
import { PublicRouter } from './PublicRouter';

export const RouterView: React.FC = () => {
  return (
    <Routes>
      <Route
        path='/'
        element={
          <PublicRouter>
            <PublicLayout />
          </PublicRouter>
        }
      >
        <Route path='' element={<Home />} />
        <Route path='login' element={<Login />} />
        <Route path='sign' element={<Sign />} />
      </Route>

      <Route
        path='/auth'
        element={
          <PrivateRouter>
            <PrivateLayout />
          </PrivateRouter>
        }
      >
        <Route path='' element={<Welcome />} />
        <Route path='create-product' element={<CreateProduct />} />
        <Route path='find-product' element={<FindorProduct />} />
      </Route>
      <Route path='*' element={<NotFound />} />
    </Routes>
  );
};
