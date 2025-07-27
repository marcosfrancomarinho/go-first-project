import { AuthUserUseCase } from '../../application/usecase/AuthUserUseCase';
import { CreatorProductUseCase } from '../../application/usecase/CreatorProductUseCase';
import { FindorProductUseCase } from '../../application/usecase/FindorProductUseCase';
import { LoginUserUseCase } from '../../application/usecase/LoginUserUseCase';
import { SignUserUseCase } from '../../application/usecase/SignUserUseCase';
import { AxiosHttpGetClient } from '../../infrastructure/http/AxiosHttpGetClient';
import { AxiosHttpPostClient } from '../../infrastructure/http/AxiosHttpPostClient';
import { CookiesClient } from '../../infrastructure/storage/CookiesClient';

export interface AppConfig {
  signUserUseCase: SignUserUseCase;
  loginUserUseCase: LoginUserUseCase;
  findorProductUseCase: FindorProductUseCase;
  creatorProductUseCase: CreatorProductUseCase;
  authUserUseCase: AuthUserUseCase;
}

export class Container {
  private static instance: Container;

  public static getInstance(): Container {
    if (!this.instance) return new Container();
    return this.instance;
  }

  public dependencies(): AppConfig {
    const baseURL: string = 'https://tool-backend-na56.onrender.com/';
    const key: string = 'user';
    const httpGetClient = new AxiosHttpGetClient(baseURL);
    const httpPostClient = new AxiosHttpPostClient(baseURL);
    const signUserUseCase = new SignUserUseCase(httpPostClient, '/register');
    const localStorageClient = new CookiesClient();
    const loginUserUseCase = new LoginUserUseCase(httpPostClient, localStorageClient, '/login', key);
    const findorProductUseCase = new FindorProductUseCase(httpGetClient, localStorageClient, '/product', key);
    const creatorProductUseCase = new CreatorProductUseCase(httpPostClient, localStorageClient, '/product', key);
    const authUserUseCase = new AuthUserUseCase(localStorageClient, key);
    return {
      signUserUseCase,
      loginUserUseCase,
      findorProductUseCase,
      creatorProductUseCase,
      authUserUseCase,
    };
  }
}
