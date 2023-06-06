import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, mergeMap, of, tap, map, EMPTY } from 'rxjs';
import { Router } from '@angular/router';
import { Store, select } from '@ngrx/store';
import { LoginState } from './login.state';
import * as LoginActions from './login.actions';
import * as LoginSelectors from './login.selectors';
import { AuthenticationService } from 'src/app/services/authentication.service';

@Injectable()
export class LoginEffects {
  initializeLogin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.initializeLogin),
      exhaustMap(() =>
        this.authenticationService.getContext()
          .pipe(
            map((user) => LoginActions.reloginRequestSuccess({ user })),
            catchError(err => {
                this.router.navigate(['/login']);
                console.error(err);
                if (err.status === 0) {
                  return of(LoginActions.reloginRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
                } else if (err.status === 500) {
                  return of(LoginActions.reloginRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }));
                }
                return EMPTY;
              }
            )
          )
      )
    )
  );

  reloginRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.reloginRequestSuccess),
      mergeMap(() => of(LoginActions.loginSuccess()))
    )
  );

  loginRequest$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.loginRequest),
      concatLatestFrom(() => this.store.select(LoginSelectors.getCredentials)),
      exhaustMap(([_, credentials]) => {
        if (credentials.email.length === 0 || credentials.password.length === 0) {
          return of(LoginActions.loginRequestFailed({ error: 'invalid credentials' }));
        }

        return this.authenticationService.login(credentials).pipe(
          map(user => LoginActions.loginRequestSuccess({ user })),
          catchError(
            (err) => { 
              if (err.status === 0) {
                return of(LoginActions.loginRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
              } else if (err.status === 400) {
                return of(LoginActions.loginRequestFailed({ error: 'Niepoprawne dane' }));
              }
              
              return of(LoginActions.loginRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }))
            }
          )
        );
      })
    )
  );

  loginRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.loginRequestSuccess),
      mergeMap(() => of(LoginActions.loginSuccess()))
    )
  );

  loginSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.loginSuccess),
      concatLatestFrom(() => this.store.pipe(select(LoginSelectors.getLoginPath))),
      tap(([_, path]) => this.router.navigate([path]))
    ), { dispatch: false }
  );

  logoutRequest$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.logoutRequest),
      exhaustMap(() => this.authenticationService.logout().pipe(
        map(() => LoginActions.logoutRequestSuccess()),
        catchError(
          (err) => {
            if (err.status === 0) {
              return of(LoginActions.logoutRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
            }

            return of(LoginActions.logoutRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }))
          }
        ))
      )
    )
  );

  logoutRequestSuccess$ = createEffect(() => 
    this.actions$.pipe(
      ofType(LoginActions.logoutRequestSuccess),
      tap(() => this.router.navigate(['/login']))
    ), { dispatch: false }
  );

  constructor(
    private actions$: Actions, 
    private router: Router, 
    private store: Store<LoginState>, 
    private authenticationService: AuthenticationService
  ) {}
}
