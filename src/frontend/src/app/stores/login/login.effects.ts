import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { loginRequest, loginRequestFailed, loginRequestSuccess, loginSuccess } from './login.actions';
import { catchError, mergeMap, of, tap, withLatestFrom } from 'rxjs';
import { Router } from '@angular/router';
import { Store, select } from '@ngrx/store';
import { LoginState } from './login.state';
import { getLoginPath } from './login.selectors';

@Injectable()
export class LoginEffects {
  login$ = createEffect(() =>
    this.actions$.pipe(
      ofType(loginRequest),
      mergeMap(() =>
        of(loginRequestSuccess({ user: {
            id: 1,
            email: 'testowy@test.com',
            deleted: null
          }})
        ).pipe(
          catchError(
            () => of(loginRequestFailed({ error: 'invalid credentials' }))
          )
        )
      )
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
