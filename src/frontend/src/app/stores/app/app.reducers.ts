import { createReducer, on } from "@ngrx/store"
import { disableHeader, enableHeader, setCurrentUrl } from "./app.actions"
import { AppState } from "./app.state"

export const initialState: AppState = {
    hideHeader: false,
    currentUrl: ''
}

export const appReducer = createReducer(
    initialState,
    on(enableHeader, (state, _) => {
        return {
            ...state,
            hideHeader: false
        }
    }),
    on(disableHeader, (state, _) => {
        return {
            ...state,
            hideHeader: true
        }
    }),
    on(setCurrentUrl, (state, action) => {
        return {
            ...state,
            currentUrl: action.currentUrl
        }
    })
)
