import { Cart } from "src/app/models/cart";

export interface CartState {
  cart: Cart;
  fetchState: FetchState;
  error: string | null;
}

export enum FetchState {
  init = 'init', loading = 'loading', success = 'success', failed = 'failed'
}
