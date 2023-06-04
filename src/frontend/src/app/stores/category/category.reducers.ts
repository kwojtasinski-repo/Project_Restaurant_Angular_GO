import { createReducer, on } from "@ngrx/store";
import { CategoryState } from "./category.state";
import { categoryAddRequestBegin, categoryAddRequestFailed, categoryFormClear, categoryFormUpdate, 
    categoryUpdateRequestBegin, categoryUpdateRequestFailed } from "./category.actions";

export const initialState: CategoryState = {
    category: null,
    error: null
}

export const categoryReducer = createReducer(
    initialState,
    on(categoryFormUpdate, (state, action) => {
        return {
            ...state,
            category: action.category
        }
    }),
    on(categoryAddRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(categoryAddRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    }),
    on(categoryFormClear, (state, _) => {
        return {
            ...state,
            product: null,
            error: null
        }
    }),
    on(categoryUpdateRequestBegin, (state, _) => {
        return {
            ...state,
            error: null
        }
    }),
    on(categoryUpdateRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error
        }
    })
);
