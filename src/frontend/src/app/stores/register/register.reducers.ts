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

export const loginReducer = createReducer(
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
            registerRequestState: RequestState.success
        }
    }),
    on(registerRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            registerRequestState: RequestState.failed
        }
    }),
);
