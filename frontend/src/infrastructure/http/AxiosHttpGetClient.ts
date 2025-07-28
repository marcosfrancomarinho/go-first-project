import axios, { AxiosError, type AxiosInstance } from 'axios';
import type { HttpGetClient } from '../../domain/gateway/HttpGetClient';
import { TokenError } from '../../shared/erros/TokenError';

export class AxiosHttpGetClient implements HttpGetClient {
  instance: AxiosInstance;
  public constructor(baseURL: string) {
    this.instance = axios.create({ baseURL });
  }

  public async get<T>(path: string, config: any): Promise<T> {
    try {
      const response = await this.instance.get(path, {
        headers: { ...config },
      });
      return response.data;
    } catch (error: any) {
      const data = error.response.data;
      if (data.code === 'TOKEN_ERROR') throw new TokenError(data.error);
      if (error instanceof AxiosError) throw new Error(data.error);
      throw new Error(error.message);
    }
  }
}
