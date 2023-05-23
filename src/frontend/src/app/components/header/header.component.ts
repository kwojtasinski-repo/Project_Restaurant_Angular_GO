import { Component, OnDestroy, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { getCurrentUrl, showHeader } from 'src/app/stores/app/app.selectors';
import { AppState } from 'src/app/stores/app/app.state';
import { getUser } from 'src/app/stores/login/login.selectors';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit, OnDestroy {
  public routerLinks: any[] = [];
  public currentUrl$ = this.store.select(getCurrentUrl);
  public showHeader$ = this.store.select(showHeader);
  public user$ = this.store.select(getUser);
  public isCollapsed = true;
  private userSubscription: Subscription = new Subscription();

  constructor(private store: Store<AppState>) { }
  
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
}
