import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { initializeLogin, loginRequest, loginRequestFailed, loginRequestSuccess, loginSuccess, logoutRequest, 
  logoutRequestFailed, logoutRequestSuccess, reloginRequestFailed, reloginRequestSuccess } from './login.actions';
import { catchError, exhaustMap, mergeMap, of, tap, map, EMPTY } from 'rxjs';
import { Router } from '@angular/router';
import { Store, select } from '@ngrx/store';
import { LoginState } from './login.state';
import { getLoginPath, selectLoginState } from './login.selectors';
import { AuthenticationService } from 'src/app/services/authentication.service';

@Injectable()
export class LoginEffects {
  initializeLogin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(initializeLogin),
      exhaustMap(() =>
        this.authenticationService.getContext()
          .pipe(
            map((user) => reloginRequestSuccess({ user })),
            catchError(err => {
                this.router.navigate(['/login']);
                console.error(err);
                if (err.status === 0) {
                  return of(reloginRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
                } else if (err.status === 500) {
                  return of(reloginRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }));
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
      ofType(reloginRequestSuccess),
      mergeMap(() => of(loginSuccess()))
    )
  )

  loginRequest$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginRequest),
      concatLatestFrom(() => this.store.select(selectLoginState)),
      exhaustMap(([_, state]) => {
        if (state.credentials.email.length === 0 || state.credentials.password.length === 0) {
          return of(loginRequestFailed({ error: 'invalid credentials' }));
        }

        return this.authenticationService.login(state.credentials).pipe(
          map(user => loginRequestSuccess({ user })),
          catchError(
            (err) => { 
              if (err.status === 0) {
                return of(loginRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
              } else if (err.status === 400) {
                return of(loginRequestFailed({ error: 'Niepoprawne dane' }));
              }
              
              return of(loginRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }))
            }
          )
        );
      })
    )
  )

  loginRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginRequestSuccess),
      mergeMap(() => of(loginSuccess()))
    )
  )

  loginSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginSuccess),
      concatLatestFrom(() => this.store.pipe(select(getLoginPath))),
      tap(([_, path]) => this.router.navigate([path]))
    ), { dispatch: false }
  )

  logoutRequest$ = createEffect(() =>
    this.actions$.pipe(
      ofType(logoutRequest),
      exhaustMap(() => this.authenticationService.logout().pipe(
        map(() => logoutRequestSuccess()),
        catchError(
          (err) => {
            if (err.status === 0) {
              return of(logoutRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
            }
                        
            return of(logoutRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }))
          }
        ))
      )
    )
  );

  logoutRequestSuccess$ = createEffect(() => 
    this.actions$.pipe(
      ofType(logoutRequestSuccess),
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
