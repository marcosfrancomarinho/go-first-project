export class Price {
  private constructor(private price: number) {}

  public static create(price: number): Price {
    this.validate(price);
    return new Price(price);
  }
  private static validate(price: number): void {
    if (!price || typeof price !== 'number' || price <= 0) {
      throw new Error('preço é obrigatorio e não poder ser menor que zero');
    }
  }
  public getValue(): number {
    return this.price;
  }
}
