export class Quantity {
  private constructor(private quantity: number) {}

  public static create(quantity: number): Quantity {
    this.validate(quantity);
    return new Quantity(quantity);
  }
  private static validate(quantity: number): void {
    if (!quantity || typeof quantity !== 'number' || quantity <= 0) {
      throw new Error('quantidade é obrigatorio e não poder ser menor que zero');
    }
  }
  public getValue(): number {
    return this.quantity;
  }
}
