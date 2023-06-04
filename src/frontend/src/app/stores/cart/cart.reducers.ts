import { createReducer, on } from "@ngrx/store";
import { CartState } from "./cart.state";
import { addProductToCartFailed, fetchCart, fetchCartFailed, fetchCartSuccess, finalizeCart, 
    finalizeCartFailed, finalizeCartSuccess, removeProductFromCartFailed } from "./cart.actions";
import { RequestState } from "src/app/models/request-state";

export const initialState: CartState = {
    cart: [],
    fetchState: RequestState.init,
    error: null,
    finalizeCartState: RequestState.init
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
    on(finalizeCart, (state, _) => {
        return {
            ...state,
            finalizeCartState: RequestState.loading
        }
    }),
    on(finalizeCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            finalizeCartState: RequestState.failed
        }
    }),
    on(finalizeCartSuccess, (state, _) => {
        return {
            ...state,
            cart: [],
            error: null,
            finalizeCartState: RequestState.success
        }
    })
);
