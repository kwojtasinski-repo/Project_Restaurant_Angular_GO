import { createFeatureSelector, createSelector } from "@ngrx/store";
import { AppState } from "./app.state";
import { appStoreName } from "./app.store.names";

export const selectLoginState = createFeatureSelector<AppState>(appStoreName);
export const showHeader = createSelector(selectLoginState, (state): boolean => state.showHeader);
export const getCurrentUrl = createSelector(selectLoginState, (state): string => state.currentUrl);
