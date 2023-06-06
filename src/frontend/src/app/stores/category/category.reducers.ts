import { createReducer, on } from "@ngrx/store";
import { CategoryState } from "./category.state";
import * as CategoryActions from "./category.actions";

export const initialState: CategoryState = {
    category: null,
    error: null
}

export const categoryReducer = createReducer(
    initialState,
    on(CategoryActions.categoryFormUpdate, (state, action) => {
        return {
            ...state,
            category: action.category
        }
    }),
    on(CategoryActions.categoryAddRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(CategoryActions.categoryAddRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(CategoryActions.categoryFormClear, (state, _) => {
        return {
            ...state,
            product: null,
            error: null
        }
    }),
    on(CategoryActions.categoryUpdateRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(CategoryActions.categoryUpdateRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    })
);
