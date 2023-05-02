import { Injectable } from '@angular/core';
import { AppState } from '../stores/app/app.state';
import { Store } from '@ngrx/store';
import { disableHeader, enableHeader, setCurrentUrl } from '../stores/app/app.actions';

@Injectable({
  providedIn: 'root'
})
export class AppService {
  constructor(private store: Store<AppState>) { }

  public showHeader() {
    this.store.dispatch(enableHeader());
  }

  public hideHeader() {
    this.store.dispatch(disableHeader());
  }

  public setCurrentUrl(url: string) {
    this.store.dispatch(setCurrentUrl({ currentUrl: url }));
  }
}
