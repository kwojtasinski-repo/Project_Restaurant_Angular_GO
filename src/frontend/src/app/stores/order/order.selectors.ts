import { createFeatureSelector, createSelector } from "@ngrx/store";
import { OrderState } from "./order.state";
import { orderStoreName } from "./order.store.names";
import { Order } from "src/app/models/order";
import { RequestState } from "src/app/models/request-state";

const selectCartState = createFeatureSelector<OrderState>(orderStoreName);
export const getOrder = createSelector(selectCartState, (state): Order | undefined => state.order);
export const getError = createSelector(selectCartState, (state): string | null => state.error);
export const getFetchState = createSelector(selectCartState, (state): RequestState => state.fetchState);
