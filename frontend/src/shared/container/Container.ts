import { AuthUserUseCase } from '../../application/usecase/AuthUserUseCase';
import { CreatorProductUseCase } from '../../application/usecase/CreatorProductUseCase';
import { FindorProductUseCase } from '../../application/usecase/FindorProductUseCase';
import { LoginUserUseCase } from '../../application/usecase/LoginUserUseCase';
import { SignUserUseCase } from '../../application/usecase/SignUserUseCase';
import { AxiosHttpClient } from '../../infrastructure/http/AxiosHttpClient';
import { LocalStorageClient } from '../../infrastructure/storage/LocalStorageClient';

export class Container {
  private static instance: Container;

  public static getInstance(): Container {
    if (!this.instance) return new Container();
    return this.instance;
  }

  public dependencies() {
    const httpClient = new AxiosHttpClient('http://localhost:8080');
    const signUserUseCase = new SignUserUseCase(httpClient);
    const localStorageClient = new LocalStorageClient();
    const loginUserUseCase = new LoginUserUseCase(httpClient, localStorageClient);
    const findorProductUseCase = new FindorProductUseCase(httpClient, localStorageClient);
    const creatorProductUseCase = new CreatorProductUseCase(httpClient, localStorageClient);
    const authUserUseCase = new AuthUserUseCase(localStorageClient)
    return {
      signUserUseCase,
      loginUserUseCase,
      findorProductUseCase,
      creatorProductUseCase,
      authUserUseCase,
    };
  }
}
