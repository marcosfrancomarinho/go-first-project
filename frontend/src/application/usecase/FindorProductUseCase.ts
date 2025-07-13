import type { HttpClient } from '../../domain/gateway/HttpClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import type { ResponseFindorProductDTO } from '../dto/ResponseFindorProductDTO';

export class FindorProductUseCase {
  public constructor(private httpClient: HttpClient, private storageClient: StorageClient) {}

  public async findAll(): Promise<ResponseFindorProductDTO[]> {
    const token: string = this.storageClient.get('token');
    const products = await this.httpClient.get<ResponseFindorProductDTO[]>('/product', { token });
    return products;
  }
}
