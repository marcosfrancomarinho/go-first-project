export class Name {
  private constructor(private name: string) {}

  public static create(name: string): Name {
    this.validate(name);
    return new Name(name.trim());
  }

  private static validate(name: string): void {
    if (!name || typeof name !== 'string' || name.trim().length === 0) {
      throw new Error('nome é obrigatorio');
    }
  }
  public getValue(): string {
    return this.name;
  }
}
