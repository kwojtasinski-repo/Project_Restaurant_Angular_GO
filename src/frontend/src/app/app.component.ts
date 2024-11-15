import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Store } from '@ngrx/store';
import { AppState } from './stores/app/app.state';
import * as AppActions from './stores/app/app.actions';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  private headerHiddenUrls: string[] = [];

  constructor(private router: Router, private appStore: Store<AppState>) { }

  public ngOnInit(): void {
    for (const route of this.router.config) {
      if (!route.data?.['hideNavBar']) {
        continue;
      }

      if (route.data?.['hideNavBar'] !== true) {
        continue;
      }

      this.headerHiddenUrls.push('/' + route.path);
    }
  }

  public onActivateRoute() {
    this.appStore.dispatch(AppActions.setCurrentUrl({ currentUrl: this.router.url }));
    if (this.headerHiddenUrls.some(url => this.getUrlWithoutParams(this.router.url) === url)) {
      this.appStore.dispatch(AppActions.disableHeader());
    } else {
      this.appStore.dispatch(AppActions.enableHeader());
    }
  }

  private getUrlWithoutParams(url: string): string {
    if (!url) {
      return '';
    }

    return url.split('?')[0];
  }
}
