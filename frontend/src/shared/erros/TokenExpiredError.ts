export class TokenExpiredError extends Error {
  public constructor(public code: string, message?: string) {
    super(message);
  }
}
