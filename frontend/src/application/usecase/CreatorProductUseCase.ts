import { Product } from '../../domain/entities/Product';
import type { HttpClient } from '../../domain/gateway/HttpClient';
import type { StorageClient } from '../../domain/gateway/StorageClient';
import { Name } from '../../domain/valuesobject/Name';
import { Price } from '../../domain/valuesobject/Price';
import { Quantity } from '../../domain/valuesobject/Quantity';
import type { RequestCreatorProductDTO } from '../dto/RequestCreatorProductDTO';
import type { ResponseCreatorProductDTO } from '../dto/ResponseCreatorProductDTO';

export class CreatorProductUseCase {
  public constructor(private httpClient: HttpClient, private storageClient: StorageClient) {}

  public async create(payload: RequestCreatorProductDTO): Promise<ResponseCreatorProductDTO> {
    const name: Name = Name.create(payload.name);
    const price: Price = Price.create(payload.price);
    const quantity: Quantity = Quantity.create(payload.quantity);
    const product: Product = Product.create(price, quantity, name);
    const { token } = this.storageClient.get<{ token: string }>(payload.key);
    const registerResponse = await this.httpClient.post<ResponseCreatorProductDTO>(
      payload.path,
      product.getProductForRegister(),
      {
        token,
      }
    );
    return registerResponse;
  }
}
