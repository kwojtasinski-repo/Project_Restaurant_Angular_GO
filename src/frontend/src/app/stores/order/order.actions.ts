import { createAction, props } from "@ngrx/store";
import { Order } from "src/app/models/order";

export const fetchOrder = createAction('[Order] Fetch Order', props<{ id: string }>());
export const fetchOrderSuccess = createAction('[Order] Fetch Order Success', props<{ order: Order | undefined }>());
export const fetchOrderFailed = createAction('[Order] Fetch Order Failed', props<{ error: string }>());

