import axios, { AxiosError, type AxiosInstance } from 'axios';
import type { HttpClient } from '../../domain/gateway/HttpClient';

export class AxiosHttpClient implements HttpClient {
  instance: AxiosInstance;
  public constructor(baseURL: string) {
    this.instance = axios.create({ baseURL });
  }
  public async post<T>(path: string, datas: any, config: any): Promise<T> {
    try {
      const response = await this.instance.post(path, datas, { headers: { ...config } });
      return response.data;
    } catch (error: any) {
      if (error instanceof AxiosError) throw new Error(error.response?.data.error);
      throw new Error(error.message);
    }
  }
  public async get<T>(path: string, config: any): Promise<T> {
    try {
      const response = await this.instance.get(path, {
        headers: { ...config },
      });
      return response.data;
    } catch (error: any) {
      if (error instanceof AxiosError) throw new Error(error.response?.data.error);
      throw new Error(error.message);
    }
  }
}
