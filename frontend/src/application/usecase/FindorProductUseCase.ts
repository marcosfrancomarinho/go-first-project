import type { HttpGetClient } from '../../domain/gateway/HttpGetClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import type { ResponseFindorProductDTO } from '../dto/ResponseFindorProductDTO';

export class FindorProductUseCase {
  public constructor(
    private httpClient: HttpGetClient,
    private storageClient: StorageClient,
    private path: string,
    private key: string
  ) {}

  public async findAll(): Promise<ResponseFindorProductDTO[]> {
    const { token } = this.storageClient.get<{ token: string }>(this.key);
    const products = await this.httpClient.get<ResponseFindorProductDTO[]>(this.path, { token });
    return products;
  }
}
