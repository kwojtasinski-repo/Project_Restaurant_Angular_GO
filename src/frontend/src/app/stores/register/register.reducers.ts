import { createReducer, on } from "@ngrx/store";
import { RegisterState } from "./register.state";
import { RequestState } from "src/app/models/request-state";
import { registerFormUpdate, registerRequestBegin, registerRequestFailed, registerRequestSuccess } from "./register.actions";

export const initialState: RegisterState = {
    email: '',
    password: '',
    passwordConfirm: '',
    error: null,
    registerRequestState: RequestState.init
}

export const registerReducer = createReducer(
    initialState,
    on(registerFormUpdate, (state, action) => {
        return {
            ...state,
            email: action.form.email,
            password: action.form.password,
            passwordConfirm: action.form.confirmPassword
        }
    }),
    on(registerRequestBegin, (state, _) => {
        return {
            ...state,
            registerRequestState: RequestState.loading
        }
    }),
    on(registerRequestSuccess, (state, _) => {
        return {
            ...state,
            registerRequestState: RequestState.success,
            email: '',
            password: '',
            passwordConfirm: '',
            error: null
        }
    }),
    on(registerRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            registerRequestState: RequestState.failed,
            email: '',
            password: '',
            passwordConfirm: '',
        }
    }),
);
