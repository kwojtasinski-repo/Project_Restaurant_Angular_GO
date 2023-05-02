import { createFeatureSelector, createSelector } from "@ngrx/store";
import { LoginState } from "./login.state";
import { loginStoreName } from "./login.store.names";
import { User } from "src/app/models/user";

export const selectLoginState = createFeatureSelector<LoginState>(loginStoreName);
export const getLoginPath = createSelector(selectLoginState, (state): string => state.path);
export const getAuthenticated = createSelector(selectLoginState, (state): boolean => state.authenticated);
export const getError = createSelector(selectLoginState, (state): any => state.error);
export const getUser = createSelector(selectLoginState, (state): User | null => state.user);
