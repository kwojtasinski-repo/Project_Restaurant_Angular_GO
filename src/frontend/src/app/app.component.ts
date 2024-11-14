import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppService } from './services/app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  private headerHiddenUrls: string[] = [];

  constructor(private router: Router, private appService: AppService) { }

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
    this.appService.setCurrentUrl(this.router.url);
    if (this.headerHiddenUrls.some(url => this.getUrlWithoutParams(this.router.url) === url)) {
        this.appService.hideHeader();
    } else {
        this.appService.showHeader();
    }
  }

  private getUrlWithoutParams(url: string): string {
    if (!url) {
      return '';
    }

    return url.split('?')[0];
  }
}
