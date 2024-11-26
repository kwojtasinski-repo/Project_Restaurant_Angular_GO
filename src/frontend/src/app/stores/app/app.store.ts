import { Injectable, Signal, signal, WritableSignal } from '@angular/core';

@Injectable({
    providedIn: 'root'
})
export class AppStore {
  private readonly _showHeader: WritableSignal<boolean> = signal(false);
  private readonly _currentUrl: WritableSignal<string> = signal('');

  get showHeader(): Signal<boolean> {
    return this._showHeader.asReadonly();
  }

  get currentUrl(): Signal<string> {
    return this._currentUrl.asReadonly();
  }

  setShowHeader(showHeader: boolean): void {
    this._showHeader.set(showHeader);
  }

  setCurrentUrl(currentUrl: string): void {
    this._currentUrl.set(currentUrl);
  }
}
