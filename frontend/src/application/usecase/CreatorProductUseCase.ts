import { Product } from '../../domain/entities/Product';
import type { HttpPostClient } from '../../domain/gateway/HttpPostClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import { Name } from '../../domain/valuesobject/Name';
import { Price } from '../../domain/valuesobject/Price';
import { Quantity } from '../../domain/valuesobject/Quantity';
import { TokenError } from '../../shared/erros/TokenError';
import type { RequestCreatorProductDTO } from '../dto/RequestCreatorProductDTO';
import type { ResponseCreatorProductDTO } from '../dto/ResponseCreatorProductDTO';

export class CreatorProductUseCase {
  public constructor(
    private httpClient: HttpPostClient,
    private storageClient: StorageClient,
    private path: string,
    private key: string
  ) {}

  public async create(payload: RequestCreatorProductDTO): Promise<ResponseCreatorProductDTO> {
    try {
      const name: Name = Name.create(payload.name);
      const price: Price = Price.create(payload.price);
      const quantity: Quantity = Quantity.create(payload.quantity);
      const product: Product = Product.create(price, quantity, name);
      const { token } = this.storageClient.get<{ token: string }>(this.key);
      const registerResponse = await this.httpClient.post<ResponseCreatorProductDTO>(this.path, product.getProductForRegister(), {
        token,
      });
      return registerResponse;
    } catch (error) {
      error instanceof TokenError && this.storageClient.delete(this.key);
      throw error;
    }
  }
}
