import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { loginRequest, loginRequestFailed, loginRequestSuccess, loginSuccess } from './login.actions';
import { catchError, mergeMap, of, tap } from 'rxjs';
import { Router } from '@angular/router';
import { Store, select } from '@ngrx/store';
import { LoginState } from './login.state';
import { getLoginPath, selectLoginState } from './login.selectors';

@Injectable()
export class LoginEffects {
  login$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginRequest),
      concatLatestFrom(() => this.store.select(selectLoginState)),
      mergeMap(([_, state]) => {
        if (state.credentials.email.length == 0 || state.credentials.password.length == 0) {
          return of(loginRequestFailed({ error: 'invalid credentials' }));
        }
        return of(loginRequestSuccess({ user: {
            id: 1,
            email: 'testowy@test.com',
            role: 'admin',
            deleted: null
          }})
        ).pipe(
          catchError(
            (err) => of(loginRequestFailed({ error: err.message }))
          )
        )
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

  constructor(private actions$: Actions, private router: Router, private store: Store<LoginState>) {}
}
