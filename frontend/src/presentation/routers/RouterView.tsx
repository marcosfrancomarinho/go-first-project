import React from 'react';
import { Route, Routes } from 'react-router-dom';
import { PublicLayout } from '../layout/PublicLayout';
import { Home } from '../pages/Home';
import { Login } from '../pages/Login';
import { Sign } from '../pages/Sign';
import { PrivateLayout } from '../layout/PrivateLayout';
import { CreateProduct } from '../pages/CreateProduct';
import { FindorProduct } from '../pages/FindorProduct';
import { Welcome } from '../pages/Welcome';

export const RouterView: React.FC = () => {
  return (
    <>
      <Routes>
        <Route path='/' element={<PublicLayout />}>
          <Route path='home' element={<Home />} />
          <Route path='login' element={<Login />} />
          <Route path='sign' element={<Sign />} />
        </Route>
        <Route path='/auth' element={<PrivateLayout />}>
          <Route path='welcome' element={<Welcome />} />
          <Route path='create-product' element={<CreateProduct />} />
          <Route path='find-product' element={<FindorProduct />} />
        </Route>
      </Routes>
    </>
  );
};
