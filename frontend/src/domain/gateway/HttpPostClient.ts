export interface HttpPostClient {
  post<T>(path:string, datas: any, config?: any): Promise<T>;
}
