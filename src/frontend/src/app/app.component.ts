import { Component, OnInit, inject } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';
import { NgxSpinnerComponent } from 'ngx-spinner';
import { FooterComponent } from './components/footer/footer.component';
import { HeaderComponent } from './components/header/header.component';
import { AppStore } from './stores/app/app.store';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
    standalone: true,
    imports: [HeaderComponent, RouterOutlet, FooterComponent, NgxSpinnerComponent]
})
export class AppComponent implements OnInit {
  private router = inject(Router);
  private appStore = inject(AppStore);

  private headerHiddenUrls: string[] = [];

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
    this.appStore.setCurrentUrl(this.router.url);
    if (this.headerHiddenUrls.some(url => this.getUrlWithoutParams(this.router.url) === url)) {
      this.appStore.disableHeader();
    } else {
      this.appStore.enableHeader();
    }
  }

  private getUrlWithoutParams(url: string): string {
    if (!url) {
      return '';
    }

    return url.split('?')[0];
  }
}
