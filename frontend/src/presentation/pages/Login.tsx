import React from 'react';
import { Link, useLocation, useNavigate } from 'react-router-dom';
import { AlertError } from '../components/AlertError';
import { Spinner } from '../components/Spinner';
import { AppContext } from '../hooks/Global';

export const Login: React.FC = () => {
  const navigate = useNavigate();
  const { loginUserUseCase } = React.useContext(AppContext)!;
  const { state } = useLocation();
  const [error, setError] = React.useState<Error | null>(null);
  const [loadding, setLoading] = React.useState<boolean>(false);
  const [email, setEmail] = React.useState<string>('');
  const [password, setPassword] = React.useState<string>('');

  const handleSubmit = async (e: any): Promise<void> => {
    try {
      e.preventDefault();
      setLoading(true);
      setError(null);
      await loginUserUseCase.login({ email, password });
      navigate('/auth');
    } catch (error) {
      setError(error as Error);
    } finally {
      setLoading(false);
    }
  };
  React.useEffect(() => {
    setEmail(state?.payload?.email ?? '');
    setPassword(state?.payload?.password ?? '');
  }, [state]);

  return (
    <div className='flex relative justify-center items-center min-h-screen bg-gray-50 px-4'>
      {error && <AlertError message={error.message} />}
      <div className='bg-white p-8 rounded-xl shadow-md w-full max-w-md'>
        <h2 className='text-2xl font-bold text-center text-blue-600 mb-6'>Login - Tools Ltda</h2>

        <form onSubmit={(e) => handleSubmit(e)} className='space-y-4'>
          <div>
            <label htmlFor='email' className='block text-sm font-medium text-gray-700'>
              E-mail
            </label>
            <input
              onChange={(e) => setEmail(e.target.value)}
              value={email}
              id='email'
              type='email'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <div>
            <label htmlFor='password' className='block text-sm font-medium text-gray-700'>
              Senha
            </label>
            <input
              onChange={(e) => setPassword(e.target.value)}
              value={password}
              id='password'
              type='password'
              className='mt-1 w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500'
            />
          </div>

          <button
            type='submit'
            className='w-full flex items-center justify-center gap-2 bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 transition'
          >
            {loadding && <Spinner />}
            Entrar
          </button>
        </form>

        <p className='text-center text-sm text-gray-600 mt-4'>
          Não tem uma conta?{' '}
          <Link to='/sign' className='text-blue-600 hover:underline'>
            Cadastre-se
          </Link>
        </p>
      </div>
    </div>
  );
};
