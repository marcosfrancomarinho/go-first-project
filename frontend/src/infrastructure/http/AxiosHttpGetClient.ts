import axios, { AxiosError, type AxiosInstance } from 'axios';
import type { HttpGetClient } from '../../domain/gateway/HttpGetClient';
import { TokenExpiredError } from '../../shared/erros/TokenExpiredError';

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
      if (error.response.data.code === 'TOKEN_ERROR')  throw new TokenExpiredError(error.response.data.code);
      if (error instanceof AxiosError) throw new Error(error.response?.data.error);
      throw new Error(error.message);
    }
  }
}
