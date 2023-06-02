import { createReducer, on } from "@ngrx/store";
import { OrderState as OrderState } from "./order.state";
import { fetchOrder, fetchOrderFailed, fetchOrderSuccess } from "./order.actions";
import { RequestState } from "src/app/models/request-state";

export const initialState: OrderState = {
    order: undefined,
    fetchState: RequestState.init,
    error: null
}

export const orderReducer = createReducer(
    initialState,
    on(fetchOrder, (state, _) => {
        return {
            ...state,
            fetchState: RequestState.loading,
            error: null
        }
    }),
    on(fetchOrderSuccess, (state, action) => {
        return {
            ...state,
            order: action.order,
            fetchState: RequestState.success,
            error: null
        }
    }),
    on(fetchOrderFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            fetchState: RequestState.failed
        }
    })
);
