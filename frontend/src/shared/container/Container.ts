import { AuthUserUseCase } from '../../application/usecase/AuthUserUseCase';
import { CreatorProductUseCase } from '../../application/usecase/CreatorProductUseCase';
import { DeleterProductUseCase } from '../../application/usecase/DeleterProductUseCase';
import { FindorProductUseCase } from '../../application/usecase/FindorProductUseCase';
import { LoginUserUseCase } from '../../application/usecase/LoginUserUseCase';
import { SignUserUseCase } from '../../application/usecase/SignUserUseCase';
import { AxiosHttpDeleteClient } from '../../infrastructure/http/AxiosHttpDeleteClient';
import { AxiosHttpGetClient } from '../../infrastructure/http/AxiosHttpGetClient';
import { AxiosHttpPostClient } from '../../infrastructure/http/AxiosHttpPostClient';
import { CookiesClient } from '../../infrastructure/storage/CookiesClient';

export interface AppConfig {
  signUserUseCase: SignUserUseCase;
  loginUserUseCase: LoginUserUseCase;
  findorProductUseCase: FindorProductUseCase;
  creatorProductUseCase: CreatorProductUseCase;
  authUserUseCase: AuthUserUseCase;
  deleterProductUseCase: DeleterProductUseCase;
}

export class Container {
  private static instance: Container;

  public static getInstance(): Container {
    if (!this.instance) return new Container();
    return this.instance;
  }

  public dependencies(): AppConfig {
    const baseURL = import.meta.env.VITE_URL
    const key: string = 'user';
    const httpGetClient = new AxiosHttpGetClient(baseURL);
    const httpPostClient = new AxiosHttpPostClient(baseURL);
    const httpDeleteClient = new AxiosHttpDeleteClient(baseURL);
    const signUserUseCase = new SignUserUseCase(httpPostClient, '/register');
    const storageClient = new CookiesClient();
    const loginUserUseCase = new LoginUserUseCase(httpPostClient, storageClient, '/login', key);
    const findorProductUseCase = new FindorProductUseCase(httpGetClient, storageClient, '/product', key);
    const creatorProductUseCase = new CreatorProductUseCase(httpPostClient, storageClient, '/product', key);
    const authUserUseCase = new AuthUserUseCase(storageClient, key);
    const deleterProductUseCase = new DeleterProductUseCase(httpDeleteClient, storageClient, '/product', key);

    return {
      signUserUseCase,
      loginUserUseCase,
      findorProductUseCase,
      creatorProductUseCase,
      authUserUseCase,
      deleterProductUseCase,
    };
  }
}
