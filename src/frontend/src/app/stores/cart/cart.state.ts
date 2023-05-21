import { Cart } from "src/app/models/cart";

export interface CartState {
  cart: Cart;
  error: string | null;
}
