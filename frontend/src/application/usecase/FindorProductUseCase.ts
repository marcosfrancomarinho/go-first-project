import type { HttpGetClient } from '../../domain/gateway/HttpGetClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import { TokenError } from '../../shared/erros/TokenError';
import type { RequestFindorProductDTO } from '../dto/RequestFindorProductDTO';
import type { ResponseFindorProductDTO } from '../dto/ResponseFindorProductDTO';

export class FindorProductUseCase {
  public constructor(
    private httpClient: HttpGetClient,
    private storageClient: StorageClient,
    private path: string,
    private key: string
  ) {}

  public async findAll(payload: RequestFindorProductDTO): Promise<ResponseFindorProductDTO> {
    try {
      const { token } = this.storageClient.get<{ token: string }>(this.key);
      const pathname: string = this.path.concat(payload.query);
      const products = await this.httpClient.get<ResponseFindorProductDTO>(pathname, { token });
      return products;
    } catch (error) {
      error instanceof TokenError && this.storageClient.delete(this.key);
      throw error;
    }
  }
}
