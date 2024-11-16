import { Product } from 'src/app/models/product';

export interface ProductState {
  product: Product | null;
  error: string | null;
}
