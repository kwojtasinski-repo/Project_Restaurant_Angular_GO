import { createReducer, on } from "@ngrx/store";
import { CartState, FetchState } from "./cart.state";
import { addProductToCartFailed, fetchCart, fetchCartFailed, fetchCartSuccess, finalizeCartFailed, finalizeCartSuccess, removeProductFromCartFailed } from "./cart.actions";

export const initialState: CartState = {
    cart: {
        products: []
    },
    fetchState: FetchState.init,
    error: null
}

export const cartReducer = createReducer(
    initialState,
    on(fetchCart, (state, _) => {
        return {
            ...state,
            fetchState: FetchState.loading,
            error: null
        }
    }),
    on(fetchCartSuccess, (state, action) => {
        return {
            ...state,
            cart: action.cart,
            fetchState: FetchState.success,
            error: null
        }
    }),
    on(fetchCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            fetchState: FetchState.failed
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
            cart: {
                products: []
            },
            error: null
        }
    })
);
