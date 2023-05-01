import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { loginRequest, loginRequestFailed, loginRequestSuccess } from './login.actions';
import { catchError, mergeMap, of } from 'rxjs';

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

  constructor(private actions$: Actions) {}
}
