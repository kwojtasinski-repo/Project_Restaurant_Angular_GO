import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AppService } from './services/app.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  constructor(private router: Router, private appService: AppService) { }

  public onActivateRoute() {
    this.appService.setCurrentUrl(this.router.url);
    if (this.router.url === '/login' || this.router.url === '/register') {
        this.appService.hideHeader();
    } else {
        this.appService.showHeader();
    }
  }
}
