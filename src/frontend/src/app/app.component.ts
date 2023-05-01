import { Component, OnInit } from '@angular/core';
import { NavigationStart, Router } from '@angular/router';
import { take } from "rxjs";
import { AuthService } from './services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  constructor(private router: Router, private authService: AuthService) { }

  ngOnInit(): void {
    this.router.events.pipe(take(1)).subscribe(event => {
      if (event instanceof NavigationStart) {
        this.authService.checkAuthenticated(event.url); 
      }
    });
  }
}
