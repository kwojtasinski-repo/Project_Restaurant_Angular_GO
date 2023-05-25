import { createReducer, on } from "@ngrx/store";
import { LoginState } from "./login.state";
import { initializeLogin, loginFormUpdate, loginRequestFailed, loginRequestSuccess, 
    logoutRequestFailed, logoutRequestSuccess } from "./login.actions";

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
    on(initializeLogin, (state, action) => {
        return {
            ...state,
            path: action.path === '' ? 'menu' : action.path
        }
    }),
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
    }),
    on(logoutRequestSuccess, (state, _) => {
        return {
            ...state,
            user: null,
            error: null,
            authenticated: false,
        }
    }),
    on(logoutRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
        }
    }),
);
