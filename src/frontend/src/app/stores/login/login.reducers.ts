import { createReducer, on } from "@ngrx/store";
import { LoginState } from "./login.state";
import { initializeLogin, loginFormUpdate, loginRequest, loginRequestFailed, loginRequestSuccess, 
    logoutRequest, 
    logoutRequestFailed, logoutRequestSuccess, reloginRequestSuccess } from "./login.actions";
import { RequestState } from "src/app/models/request-state";

export const initialState: LoginState = {
    user: null,
    authenticated: false,
    error: null,
    credentials: {
        email: '',
        password: ''
    },
    path: 'menu',
    loginRequestState: RequestState.init,
    logoutRequestState: RequestState.init
}

export const loginReducer = createReducer(
    initialState,
    on(initializeLogin, (state, action) => {
        return {
            ...state,
            path: action.path === '' ? 'menu' : action.path,
            loginRequestState: RequestState.init,
            logoutRequestState: RequestState.init
        }
    }),
    on(loginFormUpdate, (state, action) => {
        return {
            ...state,
            credentials: action.credentials
        }
    }),
    on(loginRequest, (state, _) => {
        return {
            ...state,
            error: null,
            loginRequestState: RequestState.loading
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
            },
            loginRequestState: RequestState.success
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
            },
            loginRequestState: RequestState.failed
        }
    }),
    on(logoutRequest, (state, _) => {
        return {
            ...state,
            logoutRequestState: RequestState.loading
        }
    }),
    on(logoutRequestSuccess, (state, _) => {
        return {
            ...state,
            user: null,
            error: null,
            authenticated: false,
            logoutRequestState: RequestState.success
        }
    }),
    on(logoutRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            logoutRequestState: RequestState.failed
        }
    }),
    on(reloginRequestSuccess, (state, action) => {
        return {
            ...state,
            user: action.user,
            error: null,
            authenticated: true,
            credentials: {
                email: '',
                password: ''
            },
            loginRequestState: RequestState.success
        }
    }),
);
