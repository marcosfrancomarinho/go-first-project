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
    const httpGetClient = new AxiosHttpGetClient('http://localhost:8080');
    const httpPostClient = new AxiosHttpPostClient('http://localhost:8080');
    const signUserUseCase = new SignUserUseCase(httpPostClient, '/register');
    const localStorageClient = new CookiesClient();
    const loginUserUseCase = new LoginUserUseCase(httpPostClient, localStorageClient, '/login', 'user');
    const findorProductUseCase = new FindorProductUseCase(httpGetClient, localStorageClient, '/product', 'user');
    const creatorProductUseCase = new CreatorProductUseCase(httpPostClient, localStorageClient, '/product', 'user');
    const authUserUseCase = new AuthUserUseCase(localStorageClient, 'user');
    return {
      signUserUseCase,
      loginUserUseCase,
      findorProductUseCase,
      creatorProductUseCase,
      authUserUseCase,
    };
  }
}
