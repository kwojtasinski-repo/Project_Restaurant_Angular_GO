import { Component, OnInit, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { getCurrentUrl, showHeader } from 'src/app/stores/app/app.selectors';
import { AppState } from 'src/app/stores/app/app.state';
import { getUser } from 'src/app/stores/login/login.selectors';
import { LoginState } from 'src/app/stores/login/login.state';
import { logoutRequest } from 'src/app/stores/login/login.actions';
import { RouterLink } from '@angular/router';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
    styleUrls: ['./header.component.scss'],
    standalone: true,
    imports: [CollapseModule, RouterLink, AsyncPipe]
})
export class HeaderComponent implements OnInit {
  private appStore = inject<Store<AppState>>(Store);
  private loginStore = inject<Store<LoginState>>(Store);

  public routerLinks: any[] = [];
  public currentUrl$ = this.appStore.select(getCurrentUrl);
  public showHeader$ = this.appStore.select(showHeader);
  public user$ = this.loginStore.select(getUser);
  public isCollapsed = true;
  
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

  public normalizeUrl(url: string | null): string | null {
    if (!url) {
      return url;
    }

    return url.substring(1);
  }

  public logout(): void {
    this.loginStore.dispatch(logoutRequest());
  }
}
