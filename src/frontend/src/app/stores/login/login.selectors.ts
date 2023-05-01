import { createFeatureSelector, createSelector } from "@ngrx/store";
import { LoginState } from "./login.state";
import { loginStoreName } from "./login.store.names";

export const selectLoginState = createFeatureSelector<LoginState>(loginStoreName);

export const authenticated = createSelector(selectLoginState, (state): boolean => state.authenticated);
