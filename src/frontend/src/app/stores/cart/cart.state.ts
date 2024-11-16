import { Cart } from 'src/app/models/cart';
import { RequestState } from 'src/app/models/request-state';

export interface CartState {
  cart: Cart[];
  fetchState: RequestState;
  finalizeCartState: RequestState;
  error: string | null;
}
