import { createFeatureSelector, createSelector } from "@ngrx/store";
import { CategoryState } from "./category.state";
import { categoryStoreName } from "./category.store.names";
import { Category } from "src/app/models/category";

const selectCategoryState = createFeatureSelector<CategoryState>(categoryStoreName);
export const getCategory = createSelector(selectCategoryState, (state): Category | null => state.category);
export const getError = createSelector(selectCategoryState, (state): string | null => state.error);
