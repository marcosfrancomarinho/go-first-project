import type { StorageClient } from '../../domain/gateway/StorageClient';
import Cookies from 'js-cookie';

export class CookiesClient implements StorageClient {
  public get<T>(key: string): T {
    let value: string | undefined;
    try {
      value = Cookies.get(key);
      if (!value) return '' as T;
      const json = JSON.parse(value);
      return json as T;
    } catch (error) {
      return value as T;
    }
  }
  public set(key: string, data: any): void {
    Cookies.set(key, JSON.stringify(data), { expires: 1 / 24 });
  }
  public delete(key: string): void {
    Cookies.remove(key);
  }
}
