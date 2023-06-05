import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { registerRequestBegin, registerRequestFailed, registerRequestSuccess } from './register.actions';
import { exhaustMap, of, tap, withLatestFrom } from 'rxjs';
import { RegisterState } from './register.state';
import { Store } from '@ngrx/store';
import { getForm } from './register.selectors';
import { NgxSpinnerService } from 'ngx-spinner';

@Injectable()
export class RegisterEffects {
  registerRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequestBegin),
      withLatestFrom(this.store.select(getForm)),
      tap(() => this.spinnerService.show()),
      exhaustMap((_) => of(registerRequestSuccess())
        //catchError((err) => of(registerRequestFailed(err.message))))
      )
    )
  );

  registerRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequestSuccess),
      tap(() => this.spinnerService.hide())
    ), { dispatch: false }
  );

  registerRequestFailed$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequestFailed),
      tap(() => this.spinnerService.hide())
    ), { dispatch: false }
  );

  constructor(
    private actions$: Actions,
    private store: Store<RegisterState>,
    private spinnerService: NgxSpinnerService
  ) {}
}
