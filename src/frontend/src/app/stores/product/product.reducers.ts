import { createReducer, on } from "@ngrx/store";
import { ProductState } from "./product.state";
import { productFormClear, productFormUpdate, productRequestFailed } from "./product.actions";

export const initialState: ProductState = {
    product: null,
    error: null
}

export const productReducer = createReducer(
    initialState,
    on(productFormUpdate, (state, action) => {
        return {
            ...state,
            product: action.product
        }
    }),
    on(productRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(productFormClear, (state, action) => {
        return {
            ...state,
            product: null
        }
    }),
);
