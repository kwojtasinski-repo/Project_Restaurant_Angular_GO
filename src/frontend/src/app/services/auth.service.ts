import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { LoginState } from '../stores/login/login.state';
import { Store } from '@ngrx/store';
import { take } from "rxjs";
import { getAuthenticated } from '../stores/login/login.selectors';
import { initializeLogin } from '../stores/login/login.actions';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private router: Router, private store: Store<LoginState>) {}

  public checkAuthenticated(returnUrl: string): void {
    this.store.select(getAuthenticated).pipe(take(1)).subscribe((authenticated) => {
      if (!authenticated) {
        this.store.dispatch(initializeLogin({ path: this.normalizeUrl(returnUrl) }));
        this.router.navigate(['']);
      }
    })
  }

  private normalizeUrl(url: string): string {
    let newUrl = url;

    if (url.length === 0) {
      return '';
    }

    if (url.startsWith('/')) {
      newUrl = url.substring(1);
    }

    return newUrl;
  }
}
