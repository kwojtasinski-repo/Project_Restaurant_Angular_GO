import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { User } from '../models/user';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private backendUrl = (window as any )['__env']['backendUrl'];

  constructor(private httpClient: HttpClient) { }

  public login(): Observable<User> {
    debugger
    this.httpClient.post(`${this.backendUrl}/api/sign-in`, {
      email: '',
      password: ''
    }).subscribe();

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
