import { createFeatureSelector, createSelector } from '@ngrx/store';
import { ProductState } from './product.state';
import { productStoreName } from './product.store.names';
import { Product } from 'src/app/models/product';

const selectProductState = createFeatureSelector<ProductState>(productStoreName);
export const getProduct = createSelector(selectProductState, (state): Product | null => state.product);
export const getError = createSelector(selectProductState, (state): string | null => state.error);
