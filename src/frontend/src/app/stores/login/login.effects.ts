import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { initializeLogin, loginRequest, loginRequestFailed, loginRequestSuccess, loginSuccess, logoutRequest, 
  logoutRequestFailed, logoutRequestSuccess, reloginRequestSuccess } from './login.actions';
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
            (err) => of(loginRequestFailed({ error: err.message }))
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
          (err) => of(logoutRequestFailed({ error: err.message }))
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

  constructor(private actions$: Actions, 
    private router: Router, 
    private store: Store<LoginState>, 
    private authenticationService: AuthenticationService) {}
}
