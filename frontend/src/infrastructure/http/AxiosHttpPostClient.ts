import axios, { AxiosError, type AxiosInstance } from 'axios';
import type { HttpPostClient } from '../../domain/gateway/HttpPostClient';
import { TokenError } from '../../shared/erros/TokenError';

export class AxiosHttpPostClient implements HttpPostClient {
  instance: AxiosInstance;
  public constructor(baseURL: string) {
    this.instance = axios.create({ baseURL });
  }
  public async post<T>(path: string, datas: any, config: any): Promise<T> {
    try {
      const response = await this.instance.post(path, datas, { headers: { ...config } });
      return response.data;
    } catch (error: any) {
      const data = error.response.data;
      if (data.code === 'TOKEN_ERROR') throw new TokenError(data.error);
      if (error instanceof AxiosError) throw new Error(data.error);
      throw new Error(error.message);
    }
  }
}
