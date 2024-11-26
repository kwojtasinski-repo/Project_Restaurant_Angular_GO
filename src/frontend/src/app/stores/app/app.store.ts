import { Injectable, Signal, signal, WritableSignal } from '@angular/core';

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
