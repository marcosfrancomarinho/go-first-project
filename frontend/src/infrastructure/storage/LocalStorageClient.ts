import type { StorageClient } from '../../domain/gateway/StorageClient';

export class LocalStorageClient implements StorageClient {
  public get<T>(key: string): T {
    const value: string | null = localStorage.getItem(key);
    if (!value) return '' as T;
    try {
      const json = JSON.parse(value);
      return json as T;
    } catch (error) {
      return value as T;
    }
  }
  public set(key: string, data: any): void {
    localStorage.setItem(key, JSON.stringify(data));
  }
  public delete(key: string): void {
    localStorage.removeItem(key);
  }
}
