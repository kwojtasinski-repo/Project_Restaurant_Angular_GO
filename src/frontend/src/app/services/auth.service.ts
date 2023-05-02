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
  constructor(private store: Store<LoginState>) {}

  public isAuthenticated(): boolean {
    let authenticated = false;
    this.store.select(getAuthenticated).pipe(take(1)).subscribe(a => authenticated = a);
    return authenticated;
  }
}
