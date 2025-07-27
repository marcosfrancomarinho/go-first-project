export class TokenExpiredError extends Error {
  public constructor(public code: string) {
    super();
  }
}
