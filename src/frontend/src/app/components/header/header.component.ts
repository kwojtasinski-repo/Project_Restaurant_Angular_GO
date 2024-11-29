import { Component, OnInit, effect, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { getUser } from 'src/app/stores/login/login.selectors';
import { LoginState } from 'src/app/stores/login/login.state';
import { logoutRequest } from 'src/app/stores/login/login.actions';
import { RouterLink } from '@angular/router';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { AsyncPipe } from '@angular/common';
import { AppStore } from 'src/app/stores/app/app.store';
import { BehaviorSubject } from 'rxjs';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.scss'],
    standalone: true,
    imports: [CollapseModule, RouterLink, AsyncPipe]
})
export class HeaderComponent implements OnInit {
  private readonly applicationStore = inject(AppStore);
  private readonly loginStore = inject<Store<LoginState>>(Store);
  private normalizedPath$: BehaviorSubject<string> = new BehaviorSubject('');

  public routerLinks: any[] = [];
  public currentUrl = this.applicationStore.currentUrl;
  public showHeader = this.applicationStore.showHeader;
  public user$ = this.loginStore.select(getUser);
  public isCollapsed = true;

  public constructor() {
    effect(() => {
      this.normalizedPath$.next(this.normalizeUrl(this.currentUrl()) ?? '');
    })
  }
  
  public ngOnInit(): void {
    this.routerLinks = [
      {
        name: 'Menu',
        path: 'menu'
      },
      {
        name: 'Koszyk',
        path: 'cart'
      },
      {
        name: 'Moje zamówienia',
        path: 'orders/my'
      },
      {
        name: 'Kategorie',
        path: 'categories',
        userRole: 'admin'
      }
    ]
  }

  public isActiveRoute(path: string) {
    return this.normalizedPath$.value.startsWith(path);
  }

  public logout(): void {
    this.loginStore.dispatch(logoutRequest());
  }

  private normalizeUrl(url: string | null): string | null {
    if (!url) {
      return url;
    }

    return url.substring(1);
  }
}
