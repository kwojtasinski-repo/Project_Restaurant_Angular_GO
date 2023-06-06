import { createReducer, on } from "@ngrx/store";
import { CartState } from "./cart.state";
import * as CartActions from './cart.actions';
import { RequestState } from "src/app/models/request-state";

export const initialState: CartState = {
    cart: [],
    fetchState: RequestState.init,
    error: null,
    finalizeCartState: RequestState.init
}

export const cartReducer = createReducer(
    initialState,
    on(CartActions.fetchCart, (state, _) => {
        return {
            ...state,
            fetchState: RequestState.loading,
            error: null
        }
    }),
    on(CartActions.fetchCartSuccess, (state, action) => {
        return {
            ...state,
            cart: action.cart,
            fetchState: RequestState.success,
            error: null
        }
    }),
    on(CartActions.fetchCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            fetchState: RequestState.failed
        }
    }),
    on(CartActions.addProductToCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(CartActions.removeProductFromCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(CartActions.finalizeCart, (state, _) => {
        return {
            ...state,
            finalizeCartState: RequestState.loading
        }
    }),
    on(CartActions.finalizeCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            finalizeCartState: RequestState.failed
        }
    }),
    on(CartActions.finalizeCartSuccess, (state, _) => {
        return {
            ...state,
            cart: [],
            error: null,
            finalizeCartState: RequestState.success
        }
    })
);
