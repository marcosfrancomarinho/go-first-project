import type { StorageClient } from '../../domain/gateway/StorageClient';

export class LocalStorageClient implements StorageClient {
  public get<T>(key: string): T  {
    const value = localStorage.getItem(key);
    return JSON.parse(value as string) as T;
  }
  public set(key: string, data: any): void {
    localStorage.setItem(key, JSON.stringify(data));
  }
}
