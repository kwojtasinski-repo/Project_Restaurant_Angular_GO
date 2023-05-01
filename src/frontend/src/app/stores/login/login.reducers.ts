import { createReducer, on } from "@ngrx/store";
import { LoginState } from "./login.state";
import { loginFormUpdate, loginRequestFailed, loginRequestSuccess } from "./login.actions";

export const initialState: LoginState = {
    user: null,
    authenticated: false,
    error: null,
    credentials: {
        email: '',
        password: ''
    },
    path: 'menu'
}

export const loginReducer = createReducer(
    initialState,
    on(loginFormUpdate, (state, action) => {
        return {
            ...state,
            credentials: action.credentials
        }
    }),
    on(loginRequestSuccess, (state, action) => {
        return {
            ...state,
            user: action.user,
            error: null,
            authenticated: true,
            credentials: {
                email: '',
                password: ''
            }
        }
    }),
    on(loginRequestFailed, (state, action) => {
        return {
            ...state,
            user: null,
            error: action.error,
            authenticated: false,
            credentials: {
                email: '',
                password: ''
            }
        }
    })
);
