import { createAction, props } from "@ngrx/store";
import { Cart } from "src/app/models/cart";
import { Product } from "src/app/models/product";

export const fetchCart = createAction('[Cart] Fetch Cart');
export const fetchCartSuccess = createAction('[Cart] Fetch Cart Success', props<{ cart: Cart[] }>());
export const fetchCartFailed = createAction('[Cart] Fetch Cart Failed', props<{ error: string }>());
export const addProductToCart = createAction('[Cart] Add Product to Cart', props<{ product: Product }>());
export const addProductToCartSuccess = createAction('[Cart] Add Product to Cart Success');
export const addProductToCartFailed = createAction('[Cart] Add Product to Cart Failed', props<{ error: string }>());
export const removeProductFromCart = createAction('[Cart] Remove Product from Cart', props<{ cart: Cart }>());
export const removeProductFromCartSuccess = createAction('[Cart] Remove Product from Cart Success');
export const removeProductFromCartFailed = createAction('[Cart] Remove Product from Cart Failed', props<{ error: string }>());
export const finalizeCart = createAction('[Cart] Finalize Cart');
export const finalizeCartSuccess = createAction('[Cart] Finalize Cart Success', props<{ orderId: number }>());
export const finalizeCartFailed = createAction('[Cart] Finalize Cart Failed', props<{ error: string }>());
