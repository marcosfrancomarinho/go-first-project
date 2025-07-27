import type { StorageClient } from '../../domain/gateway/StorageClient';

export class AuthUserUseCase {
  public constructor(private storageClient: StorageClient, private key:string) {}

  public isAuthenticate(): boolean {
    const user = this.storageClient.get<{ token: string }>(this.key);
    return !!user?.token;
  }
  public getNameUser(): string {
    const user = this.storageClient.get<{ name: string }>(this.key);
    if (!user?.name) throw new Error('nome do usuario não definido');
    return user.name;
  }
  public logoutUser() {
    this.storageClient.delete(this.key);
    this.isAuthenticate();
  }
}
