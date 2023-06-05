import { Injectable, Inject } from '@angular/core';
import { Observable, concatMap, EMPTY } from 'rxjs';
import { User } from '../models/user';
import { HttpClient } from '@angular/common/http';
import { Credentials } from '../models/credentials';

@Injectable({
  providedIn: 'root'
})
export class AuthenticationService {
  private signInPath = 'api/sign-in';
  private usersPath = 'api/users';
  private signOutPath = 'api/sign-out';

  constructor(private httpClient: HttpClient, @Inject('API_URL') private backendUrl: string) { }

  public login(credentials: Credentials): Observable<User> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.signInPath}`, credentials, { withCredentials: true })
      .pipe(
        concatMap(() => this.httpClient.get<User>(`${this.backendUrl}/${this.usersPath}/me`, { withCredentials: true }))
      );
  }

  public logout(): Observable<void> {
    return this.httpClient.post<void>(`${this.backendUrl}/${this.signOutPath}`, {}, { withCredentials: true });
  }

  public getContext(): Observable<User> {
    return this.httpClient.get<User>(`${this.backendUrl}/${this.usersPath}/me`, {withCredentials: true});
  }

  public register(credentials: Credentials): Observable<void> {
    return EMPTY;
  }
}
