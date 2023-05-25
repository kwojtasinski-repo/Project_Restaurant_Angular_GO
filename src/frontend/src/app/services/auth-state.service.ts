import { Injectable } from '@angular/core';
import { LoginState } from '../stores/login/login.state';
import { Store } from '@ngrx/store';
import { Observable } from "rxjs";
import { getAuthenticated, getUser } from '../stores/login/login.selectors';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class AuthStateService {
  constructor(private store: Store<LoginState>) {}

  public isAuthenticated(): Observable<boolean> {
    return this.store.select(getAuthenticated);
  }

  public getUser(): Observable<User | null> {
    return this.store.select(getUser);
  }
}
