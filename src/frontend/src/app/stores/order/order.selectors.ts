import { createFeatureSelector, createSelector } from "@ngrx/store";
import { OrderState, FetchState } from "./order.state";
import { orderStoreName } from "./order.store.names";
import { Order } from "src/app/models/order";

export const selectCartState = createFeatureSelector<OrderState>(orderStoreName);
export const getOrder = createSelector(selectCartState, (state): Order | undefined => state.order);
export const getError = createSelector(selectCartState, (state): string | null => state.error);
export const getFetchState = createSelector(selectCartState, (state): FetchState => state.fetchState);
