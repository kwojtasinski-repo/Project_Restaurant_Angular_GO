import { createFeatureSelector, createSelector } from '@ngrx/store';
import { CartState } from './cart.state';
import { cartStoreName } from './cart.store.names';
import { Cart } from 'src/app/models/cart';
import { RequestState } from 'src/app/models/request-state';

const selectCartState = createFeatureSelector<CartState>(cartStoreName);
export const getCart = createSelector(selectCartState, (state): Cart[] => state.cart ?? []);
export const getError = createSelector(selectCartState, (state): string | null => state?.error);
export const getFetchState = createSelector(selectCartState, (state): RequestState => state.fetchState);
export const getFinalizeState = createSelector(selectCartState, (state): RequestState => state.finalizeCartState);
