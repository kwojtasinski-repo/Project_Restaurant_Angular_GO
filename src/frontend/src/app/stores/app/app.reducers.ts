import { createReducer, on } from "@ngrx/store"
import { disableHeader, enableHeader, setCurrentUrl } from "./app.actions"
import { AppState } from "./app.state"

export const initialState: AppState = {
    showHeader: false,
    currentUrl: ''
}

export const appReducer = createReducer(
    initialState,
    on(enableHeader, (state, _) => {
        return {
            ...state,
            showHeader: true
        }
    }),
    on(disableHeader, (state, _) => {
        return {
            ...state,
            showHeader: false
        }
    }),
    on(setCurrentUrl, (state, action) => {
        return {
            ...state,
            currentUrl: action.currentUrl
        }
    })
)
