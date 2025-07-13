export class Password {
  private constructor(private password: string) {}

  public static create(password: string): Password {
    this.validate(password);
    return new Password(password.trim());
  }
  private static validate(password: string): void {
    if (!password || typeof password !== 'string' || password.trim().length === 0) {
      throw new Error('senha é obrigatorio');
    }
  }
  public getValue(): string {
    return this.password;
  }
}
