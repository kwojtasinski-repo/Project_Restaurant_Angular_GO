import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AppService } from './services/app.service';

let headerHiddenUrls: string[] = [];

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  constructor(private router: Router, private appService: AppService) { }

  public ngOnInit(): void {
    for (let route of this.router.config) {
      if (!route.data?.['hideNavBar']) {
        continue;
      }

      if (route.data?.['hideNavBar'] !== true) {
        continue;
      }

      headerHiddenUrls.push('/' + route.path);
    }
  }

  public onActivateRoute() {
    this.appService.setCurrentUrl(this.router.url);
    if (headerHiddenUrls.some(url => this.getUrlWithoutParams(this.router.url) === url)) {
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
