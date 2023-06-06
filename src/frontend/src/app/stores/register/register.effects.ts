import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { registerRequestBegin, registerRequestFailed, registerRequestSuccess } from './register.actions';
import { catchError, exhaustMap, map, of, tap, withLatestFrom } from 'rxjs';
import { RegisterState } from './register.state';
import { Store } from '@ngrx/store';
import { getForm } from './register.selectors';
import { NgxSpinnerService } from 'ngx-spinner';
import { Router } from '@angular/router';
import { AuthenticationService } from 'src/app/services/authentication.service';

@Injectable()
export class RegisterEffects {
  registerRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequestBegin),
      withLatestFrom(this.store.select(getForm)),
      tap(() => this.spinnerService.show()),
      exhaustMap(([_, form]) => this.authenticationService.register({
          email: form.email,
          password: form.password
        })
        .pipe(
          map(() => registerRequestSuccess()),
          catchError((err) => of(registerRequestFailed(err.message)))
        )
      )
    )
  );

  registerRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(registerRequestSuccess),
      tap(() => {
        this.spinnerService.hide();
        this.router.navigate(['/register-success'], { queryParams: { registerState: 'success' }});
      })
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
    private spinnerService: NgxSpinnerService,
    private authenticationService: AuthenticationService,
    private router: Router
  ) {}
}
