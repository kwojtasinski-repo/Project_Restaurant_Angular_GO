import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { getCurrentUrl, showHeader } from 'src/app/stores/app/app.selectors';
import { AppState } from 'src/app/stores/app/app.state';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {
  public routerLinks: any[] = [];
  public currentUrl$ = this.store.select(getCurrentUrl);
  public showHeader$ = this.store.select(showHeader);

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
