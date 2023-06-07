import { createReducer, on } from "@ngrx/store";
import { RegisterState } from "./register.state";
import { RequestState } from "src/app/models/request-state";
import * as RegisterActions from "./register.actions";

export const initialState: RegisterState = {
    email: '',
    password: '',
    passwordConfirm: '',
    error: null,
    registerRequestState: RequestState.init
}

export const registerReducer = createReducer(
    initialState,
    on(RegisterActions.registerFormUpdate, (state, action) => {
        return {
            ...state,
            email: action.form.email,
            password: action.form.password,
            passwordConfirm: action.form.confirmPassword
        }
    }),
    on(RegisterActions.registerRequestBegin, (state, _) => {
        return {
            ...state,
            registerRequestState: RequestState.loading
        }
    }),
    on(RegisterActions.registerRequestSuccess, (state, _) => {
        return {
            ...state,
            registerRequestState: RequestState.success,
            email: '',
            password: '',
            passwordConfirm: '',
            error: null
        }
    }),
    on(RegisterActions.registerRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            registerRequestState: RequestState.failed,
            email: '',
            password: '',
            passwordConfirm: '',
        }
    }),
    on(RegisterActions.clearErrors, (state, _) => {
        return {
            ...state,
            error: null
        }
    })
);
