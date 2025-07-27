import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { PublicLayout } from '../../ui/layout/PublicLayout';
import { Home } from '../../ui/pages/Home';
import { Login } from '../../ui/pages/Login';
import { Sign } from '../../ui/pages/Sign';
import { PrivateLayout } from '../../ui/layout/PrivateLayout';
import { CreateProduct } from '../../ui/pages/CreateProduct';
import { FindorProduct } from '../../ui/pages/FindorProduct';
import { Welcome } from '../../ui/pages/Welcome';
import { PrivateRouter } from './PrivateRouter';
import { PublicRouter } from './PublicRouter';
import { NotFound } from '../../ui/pages/NotFound';

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
