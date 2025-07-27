import type { StorageClient } from '../../domain/gateway/StorageClient';
import cookies from 'js-cookie';

export class CookiesClient implements StorageClient {
  public get<T>(key: string): T {
    const value = cookies.get(key);
    return value ? JSON.parse(value as string) : ('' as T);
  }
  public set(key: string, data: any): void {
    cookies.set(key, JSON.stringify(data), { expires: 1 / 24 });
  }
  public delete(key: string): void {
    cookies.remove(key);
  }
}
