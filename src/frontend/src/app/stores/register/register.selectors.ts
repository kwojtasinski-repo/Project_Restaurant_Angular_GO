import { createFeatureSelector, createSelector } from "@ngrx/store";
import { RegisterState } from "./register.state";
import { registerStoreName } from "./register.store.names";
import { RequestState } from "src/app/models/request-state";

const selectRegisterState = createFeatureSelector<RegisterState>(registerStoreName);
export const getEmail = createSelector(selectRegisterState, (state): string => state.email);
export const getPassword = createSelector(selectRegisterState, (state): string => state.password);
export const getPasswordConfirm = createSelector(selectRegisterState, (state): string => state.passwordConfirm);
export const getForm = createSelector(selectRegisterState, (state) => ({
    email: state.email,
    password: state.password,
    passwordConfirm: state.passwordConfirm,
}));
export const getRegisterRequestState = createSelector(selectRegisterState, (state): RequestState => state.registerRequestState);
export const getError = createSelector(selectRegisterState, (state): any => state.error);
