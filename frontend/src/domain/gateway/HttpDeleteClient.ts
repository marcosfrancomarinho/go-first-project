export interface HttpDeleteClient {
  delete<T>(path:string, datas: any, config?: any): Promise<T>;
}
