import { createReducer, on } from '@ngrx/store';
import { ProductState } from './product.state';
import * as ProductActions from './product.actions';

export const initialState: ProductState = {
    product: null,
    error: null
}

export const productReducer = createReducer(
    initialState,
    on(ProductActions.productFormUpdate, (state, action) => {
        return {
            ...state,
            product: action.product
        }
    }),
    on(ProductActions.productAddRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(ProductActions.productAddRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(ProductActions.productFormClear, (state, _) => {
        return {
            ...state,
            product: null,
            error: null
        }
    }),
    on(ProductActions.productUpdateRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(ProductActions.productUpdateRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(ProductActions.clearErrors, (state, _) => {
        return {
            ...state,
            error: null
        }
    })
);
