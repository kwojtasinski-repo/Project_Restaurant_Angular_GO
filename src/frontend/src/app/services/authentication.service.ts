import { Injectable } from '@angular/core';
import { Observable, catchError, concatMap, exhaustMap, from, map, of, take } from 'rxjs';
import { User } from '../models/user';
import { HttpClient } from '@angular/common/http';
import { Credentials } from '../models/credentials';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private backendUrl = (window as any )['__env']['backendUrl'];

  constructor(private httpClient: HttpClient) { }

  public login(credentials: Credentials): Observable<User> {
    return this.httpClient.post<void>(`${this.backendUrl}/api/sign-in`, credentials, { withCredentials: true })
      .pipe(
        concatMap(() => this.httpClient.get<User>(`${this.backendUrl}/api/users/me`, { withCredentials: true }))
      );
  }

  public logout(): Observable<void> {
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }
}
