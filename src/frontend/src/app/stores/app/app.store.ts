import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';

type ApplicationState = {
  showHeader: boolean;
  currentUrl: string;
};

const initialState: ApplicationState = {
  showHeader: false,
  currentUrl: ''
};

export const AppStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store) => ({
    enableHeader(): void {
      patchState(store, (state) => ({
        ...state,
        showHeader: true
      }))
    },
    disableHeader(): void  {
      patchState(store, (state) => ({
        ...state,
        showHeader: false
      }))
    },
    setCurrentUrl(currentUrl: string): void {
      patchState(store, (state) => ({
        ...state,
        currentUrl
      }))
    }
  }))
);
