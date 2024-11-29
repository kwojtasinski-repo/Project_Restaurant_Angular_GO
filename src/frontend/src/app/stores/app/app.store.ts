import { Injectable, Signal, signal, WritableSignal } from '@angular/core';
import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';

@Injectable({
    providedIn: 'root'
})
export class AppStore {
  private readonly _showHeader: WritableSignal<boolean> = signal(false);
  private readonly _currentUrl: WritableSignal<string> = signal('');

  public get showHeader(): Signal<boolean> {
    return this._showHeader.asReadonly();
  }

  public get currentUrl(): Signal<string> {
    return this._currentUrl.asReadonly();
  }

  public enableHeader(): void {
    this._showHeader.set(true);
  }

  public disableHeader(): void {
    this._showHeader.set(false);
  }

  public setCurrentUrl(currentUrl: string): void {
    this._currentUrl.set(currentUrl);
  }
}

type ApplicationState = {
  showHeader: boolean;
  currentUrl: string;
};

const initialState: ApplicationState = {
  showHeader: false,
  currentUrl: ''
};

export const ApplicationStore = signalStore(
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

