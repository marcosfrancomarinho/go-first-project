export interface HttpClient {
  post<T>(path: string, datas: any, config?: any): Promise<T>;
  get<T>(path: string, config: any): Promise<T>;
}
