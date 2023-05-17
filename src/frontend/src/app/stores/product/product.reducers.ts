import { createReducer, on } from "@ngrx/store";
import { ProductState } from "./product.state";
import { productFormClear, productFormUpdate, productAddRequestFailed, productUpdateRequestFailed, productAddRequestBegin, productUpdateRequestBegin } from "./product.actions";

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
    on(productAddRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(productAddRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(productFormClear, (state, _) => {
        return {
            ...state,
            product: null,
            error: null
        }
    }),
    on(productUpdateRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(productUpdateRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    })
);
