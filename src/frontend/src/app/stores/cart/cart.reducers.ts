import { createReducer, on } from "@ngrx/store";
import { CartState } from "./cart.state";
import { addProductToCartFailed, fetchCart, fetchCartFailed, fetchCartSuccess, finalizeCartFailed, finalizeCartSuccess, removeProductFromCartFailed } from "./cart.actions";
import { RequestState } from "src/app/models/request-state";

export const initialState: CartState = {
    cart: [],
    fetchState: RequestState.init,
    error: null
}

export const cartReducer = createReducer(
    initialState,
    on(fetchCart, (state, _) => {
        return {
            ...state,
            fetchState: RequestState.loading,
            error: null
        }
    }),
    on(fetchCartSuccess, (state, action) => {
        return {
            ...state,
            cart: action.cart,
            fetchState: RequestState.success,
            error: null
        }
    }),
    on(fetchCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            fetchState: RequestState.failed
        }
    }),
    on(addProductToCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(removeProductFromCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(finalizeCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(finalizeCartSuccess, (state, _) => {
        return {
            ...state,
            cart: [],
            error: null
        }
    })
);
