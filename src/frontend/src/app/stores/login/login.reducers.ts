import { createReducer, on } from "@ngrx/store";
import { LoginState } from "./login.state";
import { loginRequestFailed, loginRequestSuccess } from "./login.actions";

export const initialState: LoginState = {
    user: null,
    authenticated: false,
    error: null
}

export const loginReducer = createReducer(
    initialState,
    on(loginRequestSuccess, (state, action) => {
        return {
            ...state,
            user: action.user,
            error: null,
            authenticated: true
        }
    }),
    on(loginRequestFailed, (state, action) => {
        return {
            ...state,
            user: null,
            error: action.error,
            authenticated: false
        }
    })
);
