import type { StorageClient } from '../../domain/gateway/StorageClient';

export class LocalStorageClient implements StorageClient {
  public get<T>(key: string): T {
    this.keyValidator(key);
    const value = localStorage.getItem(key);
    if (!value) throw new Error(`A chave ${key} não tem valor definido`);
    return JSON.parse(value) as T;
  }
  public set(key: string, data: any): void {
    this.keyValidator(key);
    localStorage.setItem(key, JSON.stringify(data));
  }

  private keyValidator(key: string): void {
    if (!key || typeof key !== 'string' || key.length === 0) throw new Error('chave nao foi informada');
  }
}
