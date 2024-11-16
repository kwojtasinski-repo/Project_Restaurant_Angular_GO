import { createReducer, on } from '@ngrx/store';
import { OrderState as OrderState } from './order.state';
import * as OrderActions from './order.actions';
import { RequestState } from 'src/app/models/request-state';

export const initialState: OrderState = {
    order: undefined,
    fetchState: RequestState.init,
    error: null
}

export const orderReducer = createReducer(
    initialState,
    on(OrderActions.fetchOrder, (state, _) => {
        return {
            ...state,
            fetchState: RequestState.loading,
            error: null
        }
    }),
    on(OrderActions.fetchOrderSuccess, (state, action) => {
        return {
            ...state,
            order: action.order,
            fetchState: RequestState.success,
            error: null
        }
    }),
    on(OrderActions.fetchOrderFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            fetchState: RequestState.failed
        }
    })
);
