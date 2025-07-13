export class Email {
  private constructor(private email: string) {}

  public static create(email: string): Email {
    this.validate(email);
    return new Email(email.trim());
  }
  private static validate(email: string): void {
    if (!email || typeof email !== 'string' || email.trim().length === 0) {
      throw new Error('senha é obrigatorio');
    }
  }
  public getValue(): string {
    return this.email;
  }
}
