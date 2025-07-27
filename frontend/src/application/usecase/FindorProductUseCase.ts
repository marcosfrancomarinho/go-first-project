import type { HttpClient } from '../../domain/gateway/HttpClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import type { ResponseFindorProductDTO } from '../dto/ResponseFindorProductDTO';
import type { ResquestFindorProductDTO } from '../dto/ResquestFindorProductDTO';

export class FindorProductUseCase {
  public constructor(private httpClient: HttpClient, private storageClient: StorageClient) {}

  public async findAll(payload: ResquestFindorProductDTO): Promise<ResponseFindorProductDTO[]> {
    const { token } = this.storageClient.get<{ token: string }>(payload.key);
    const products = await this.httpClient.get<ResponseFindorProductDTO[]>(payload.path, { token });
    return products;
  }
}
