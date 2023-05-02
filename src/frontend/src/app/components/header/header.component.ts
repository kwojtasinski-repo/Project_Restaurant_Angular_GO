import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { customRoutes } from 'src/app/routes';
import { getCurrentUrl, hideHeader } from 'src/app/stores/app/app.selectors';
import { AppState } from 'src/app/stores/app/app.state';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  public routerLinks: any[] = [];
  public currentUrl$ = this.store.select(getCurrentUrl);
  public hideHeader$ = this.store.select(hideHeader);

  constructor(private store: Store<AppState>) { }
  
  public ngOnInit(): void {
    this.routerLinks = [
      {
        name: 'Menu',
        path: 'menu'
      }
    ]
  }
}
