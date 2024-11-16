import { createReducer, on } from '@ngrx/store'
import * as AppActions from './app.actions'
import { AppState } from './app.state'

export const initialState: AppState = {
    showHeader: false,
    currentUrl: ''
}

export const appReducer = createReducer(
    initialState,
    on(AppActions.enableHeader, (state, _) => {
        return {
            ...state,
            showHeader: true
        }
    }),
    on(AppActions.disableHeader, (state, _) => {
        return {
            ...state,
            showHeader: false
        }
    }),
    on(AppActions.setCurrentUrl, (state, action) => {
        return {
            ...state,
            currentUrl: action.currentUrl
        }
    })
)
