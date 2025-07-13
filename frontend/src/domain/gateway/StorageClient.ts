export interface StorageClient {
  set(key: string, data: any): void;
  get<T>(key: string): T;
}
