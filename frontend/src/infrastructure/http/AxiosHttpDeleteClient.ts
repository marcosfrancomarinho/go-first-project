import axios, { AxiosError, type AxiosInstance } from 'axios';
import type { HttpDeleteClient } from '../../domain/gateway/HttpDeleteClient';
import { TokenExpiredError } from '../../shared/erros/TokenExpiredError';

export class AxiosHttpDeleteClient implements HttpDeleteClient {
  instance: AxiosInstance;
  public constructor(baseURL: string) {
    this.instance = axios.create({ baseURL });
  }

  public async delete<T>(path: string, datas: any, config: any): Promise<T> {
    try {
      const response = await this.instance.delete(`${path}/${datas}`, { headers: { ...config } });
      return response.data;
    } catch (error: any) {
      const data = error.response.data;
      if (data.code === 'TOKEN_ERROR') throw new TokenExpiredError(data.code, data.error);
      if (error instanceof AxiosError) throw new Error(error.response?.data.error);
      throw new Error(error.message);
    }
  }
}
