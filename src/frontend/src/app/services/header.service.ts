import { Injectable } from '@angular/core';
import { AppState } from '../stores/app/app.state';
import { Store } from '@ngrx/store';
import { customRoutes } from '../routes';
import { disableHeader, enableHeader, setCurrentUrl } from '../stores/app/app.actions';

@Injectable({
  providedIn: 'root'
})
export class HeaderService {
  constructor(private store: Store<AppState>) { }

  public checkHeader(url: string) {
    
  }

  public showHeader() {
    this.store.dispatch(enableHeader());
  }

  public hideHeader() {
    this.store.dispatch(disableHeader());
  }

  private normalizeUrl(url: string): string {
    if (!url) {
      return url;
    }

    if (url.length === 0) {
      return url;
    }

    if (url.startsWith('/')) {
      return url.substring(1);
    }

    return url;
  }
}
