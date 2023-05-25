import { Component, OnDestroy, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { getCurrentUrl, showHeader } from 'src/app/stores/app/app.selectors';
import { AppState } from 'src/app/stores/app/app.state';
import { getUser } from 'src/app/stores/login/login.selectors';
import { Subscription } from 'rxjs';
import { LoginState } from 'src/app/stores/login/login.state';
import { logoutRequest } from 'src/app/stores/login/login.actions';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit, OnDestroy {
  public routerLinks: any[] = [];
  public currentUrl$ = this.appStore.select(getCurrentUrl);
  public showHeader$ = this.appStore.select(showHeader);
  public user$ = this.loginStore.select(getUser);
  public isCollapsed = true;
  private userSubscription: Subscription = new Subscription();

  constructor(private appStore: Store<AppState>, private loginStore: Store<LoginState>) { }
  
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
      }
    ]

    this.userSubscription = this.user$
      .subscribe(user => {
        if (user?.role === 'admin') {
          this.routerLinks.push({
            name: 'Kategorie',
            path: 'categories'
          })
        }
      })
  }

  public ngOnDestroy(): void {
    this.userSubscription.unsubscribe();
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
