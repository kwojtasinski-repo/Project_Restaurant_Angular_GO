import { createFeatureSelector, createSelector } from "@ngrx/store";
import { ProductState } from "./category.state";
import { productStoreName } from "./category.store.names";
import { Category } from "src/app/models/category";

export const selectCategoryState = createFeatureSelector<ProductState>(productStoreName);
export const getCategory = createSelector(selectCategoryState, (state): Category | null => state.category);
export const getError = createSelector(selectCategoryState, (state): string | null => state.error);
