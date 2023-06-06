import { createReducer, on } from "@ngrx/store";
import { LoginState } from "./login.state";
import * as CategoryActions from "./login.actions";
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
    on(CategoryActions.initializeLogin, (state, action) => {
        return {
            ...state,
            path: action.path === '' ? 'menu' : action.path,
            loginRequestState: RequestState.init,
            logoutRequestState: RequestState.init
        }
    }),
    on(CategoryActions.loginFormUpdate, (state, action) => {
        return {
            ...state,
            credentials: action.credentials
        }
    }),
    on(CategoryActions.loginRequest, (state, _) => {
        return {
            ...state,
            error: null,
            loginRequestState: RequestState.loading
        }
    }),
    on(CategoryActions.loginRequestSuccess, (state, action) => {
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
    on(CategoryActions.loginRequestFailed, (state, action) => {
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
    on(CategoryActions.logoutRequest, (state, _) => {
        return {
            ...state,
            logoutRequestState: RequestState.loading
        }
    }),
    on(CategoryActions.logoutRequestSuccess, (state, _) => {
        return {
            ...state,
            user: null,
            error: null,
            authenticated: false,
            logoutRequestState: RequestState.success
        }
    }),
    on(CategoryActions.logoutRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            logoutRequestState: RequestState.failed
        }
    }),
    on(CategoryActions.reloginRequestSuccess, (state, action) => {
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
    on(CategoryActions.reloginRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            loginRequestState: RequestState.failed
        }
    })
);
