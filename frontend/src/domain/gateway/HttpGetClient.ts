export interface HttpGetClient {
  get<T>(path:string, config: any): Promise<T>;
}
