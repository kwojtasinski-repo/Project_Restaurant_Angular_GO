import { createReducer, on } from '@ngrx/store';
import { LoginState } from './login.state';
import * as LoginActions from './login.actions';
import { RequestState } from 'src/app/models/request-state';

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
    on(LoginActions.loginFormUpdate, (state, action) => {
        return {
            ...state,
            credentials: action.credentials
        }
    }),
    on(LoginActions.loginRequest, (state, _) => {
        return {
            ...state,
            error: null,
            loginRequestState: RequestState.loading
        }
    }),
    on(LoginActions.loginRequestSuccess, (state, action) => {
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
    on(LoginActions.loginRequestFailed, (state, action) => {
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
    on(LoginActions.logoutRequest, (state, _) => {
        return {
            ...state,
            logoutRequestState: RequestState.loading
        }
    }),
    on(LoginActions.logoutRequestSuccess, (state, _) => {
        return {
            ...state,
            user: null,
            error: null,
            authenticated: false,
            logoutRequestState: RequestState.success
        }
    }),
    on(LoginActions.logoutRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            logoutRequestState: RequestState.failed
        }
    }),
    on(LoginActions.reloginRequestSuccess, (state, action) => {
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
    on(LoginActions.reloginRequestFailed, (state, action) => {
        return {
            ...state,
            error: action.error,
            loginRequestState: RequestState.failed
        }
    }),
    on(LoginActions.setTargetPath, (state, action) => ({
        ...state,
        path: (!action.path || action.path  === '') ? 'menu' : action.path
    }))
);
