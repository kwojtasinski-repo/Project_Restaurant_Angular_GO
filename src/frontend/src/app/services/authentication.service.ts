import { Injectable, Inject } from '@angular/core';
import { Observable, concatMap } from 'rxjs';
import { User } from '../models/user';
import { HttpClient } from '@angular/common/http';
import { Credentials } from '../models/credentials';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  constructor(private httpClient: HttpClient, @Inject('API_URL') private backendUrl: string) { }

  public login(credentials: Credentials): Observable<User> {
    return this.httpClient.post<void>(`${this.backendUrl}/api/sign-in`, credentials, { withCredentials: true })
      .pipe(
        concatMap(() => this.httpClient.get<User>(`${this.backendUrl}/api/users/me`, { withCredentials: true }))
      );
  }

  public logout(): Observable<void> {
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }

  public getContext(): Observable<User> {
    return this.httpClient.get<User>(`${this.backendUrl}/api/users/me`, {withCredentials: true});
  }
}
