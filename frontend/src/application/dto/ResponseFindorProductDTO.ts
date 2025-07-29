export interface AllProducts {
  id: string;
  name: string;
  price: number;
  quantity: number;
  total: number;
}

export interface ResponseFindorProductDTO {
  all: AllProducts[];
  total: number;
}
