import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {

  constructor() { }

  public login(): Observable<User> {
    return of({
        id: 1,
        email: 'testowy@test.com',
        role: 'admin',
        deleted: null
      }
    );
  }

  public logout(): Observable<void> {
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }
}
