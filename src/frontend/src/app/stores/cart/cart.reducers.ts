import { createReducer, on } from "@ngrx/store";
import { CartState } from "./cart.state";
import { addProductToCartFailed, fetchCartFailed, fetchCartSuccess, finalizeCartFailed, finalizeCartSuccess, removeProductFromCartFailed } from "./cart.actions";

export const initialState: CartState = {
    cart: {
        products: []
    },
    error: null
}

export const cartReducer = createReducer(
    initialState,
    on(fetchCartSuccess, (state, action) => {
        return {
            ...state,
            cart: action.cart
        }
    }),
    on(fetchCartFailed, (state, action) => {
        return {
            ...state,
            error: action.error
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
            }
        }
    })
);
