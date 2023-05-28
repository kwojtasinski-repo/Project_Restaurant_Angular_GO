import { Injectable } from '@angular/core';
import { Observable, catchError, map, of, take } from 'rxjs';
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
    debugger
    this.httpClient.post(`${this.backendUrl}/api/sign-in`, credentials, { withCredentials: true })
      .pipe(take(1))
      .subscribe(() => {
        debugger;
      });
    return this.httpClient.post<User>(`${this.backendUrl}/api/sign-in`, credentials);
  }

  public logout(): Observable<void> {
    return new Observable((ob) => { ob.next(); ob.complete(); });
  }
}
