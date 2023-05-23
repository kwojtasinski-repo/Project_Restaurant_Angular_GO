import { createReducer, on } from "@ngrx/store";
import { OrderState as OrderState, FetchState } from "./order.state";
import { fetchOrder, fetchOrderFailed, fetchOrderSuccess } from "./order.actions";

export const initialState: OrderState = {
    order: undefined,
    fetchState: FetchState.init,
    error: null
}

export const orderReducer = createReducer(
    initialState,
    on(fetchOrder, (state, _) => {
        return {
            ...state,
            fetchState: FetchState.loading,
            error: null
        }
    }),
    on(fetchOrderSuccess, (state, action) => {
        return {
            ...state,
            order: action.order,
            fetchState: FetchState.success,
            error: null
        }
    }),
    on(fetchOrderFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            fetchState: FetchState.failed
        }
    })
);
