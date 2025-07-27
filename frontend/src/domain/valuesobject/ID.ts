export class ID {
  private constructor(private id: string) {}

  public static create(id: string): ID {
    this.validate(id);
    return new ID(id.trim());
  }

  private static validate(id: string): void {
    if (!id || typeof id !== 'string' || id.trim().length === 0) {
      throw new Error('Id é obrigatorio');
    }
  }
  public getValue(): string {
    return this.id;
  }
}
