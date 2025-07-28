import type { HttpDeleteClient } from '../../domain/gateway/HttpDeleteClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import { ID } from '../../domain/valuesobject/ID';
import { TokenError } from '../../shared/erros/TokenError';
import type { RequestDeleterProductDTO } from '../dto/RequestDeleterProductDTO';
import type { ResponseDeleterProductDTO } from '../dto/ResponseDeleterProductDTO';

export class DeleterProductUseCase {
  public constructor(
    private httpClient: HttpDeleteClient,
    private storageClient: StorageClient,
    private path: string,
    private key: string
  ) {}

  public async delete(payload: RequestDeleterProductDTO): Promise<ResponseDeleterProductDTO> {
    try {
      const id: ID = ID.create(payload.id);
      const { token } = this.storageClient.get<{ token: string }>(this.key);
      const exclusionResponse: ResponseDeleterProductDTO = await this.httpClient.delete(this.path, id.getValue(), { token });
      return exclusionResponse;
    } catch (error) {
      error instanceof TokenError && this.storageClient.delete(this.key);
      throw error;
    }
  }
}
