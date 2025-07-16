import type { StorageClient } from '../../domain/gateway/StorageClient';

export class AuthUserUseCase {
  public constructor(private storageClient: StorageClient) {}

  public isAuthenticate(): boolean {
    const user = this.storageClient.get<{ token: string }>('user');
    return !!user?.token;
  }
  public getNameUser(): string {
    const user = this.storageClient.get<{ name: string }>('user');
    if (!user?.name) throw new Error('nome do usuario não definido');
    return user.name;
  }
}
