import type { Name } from '../valuesobject/Name';
import type { Price } from '../valuesobject/Price';
import type { Quantity } from '../valuesobject/Quantity';

export class Product {
  private constructor(private price: Price, private quantity: Quantity, private name: Name) {}

  public static create(price: Price, quantity: Quantity, name: Name): Product {
    return new Product(price, quantity, name);
  }

  public getProductForRegister() {
    return {
      name: this.name.getValue(),
      price: this.price.getValue(),
      quantity: this.quantity.getValue(),
    };
  }
}
