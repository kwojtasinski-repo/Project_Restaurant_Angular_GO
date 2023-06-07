import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import * as RegisterActions from './register.actions';
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
      ofType(RegisterActions.registerRequestBegin),
      withLatestFrom(this.store.select(getForm)),
      tap(() => this.spinnerService.show()),
      exhaustMap(([_, form]) => this.authenticationService.register({
          email: form.email,
          password: form.password
        })
        .pipe(
          map(() => RegisterActions.registerRequestSuccess()),
          catchError((err) => {
            console.error(err);
            if (err.status === 0) {
              return of(RegisterActions.registerRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
            } else if (err.status >= 500) {
              of(RegisterActions.registerRequestFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
            }

            return of(RegisterActions.registerRequestFailed({ error: 'Niepoprawne hasło lub email' }));
          })
        )
      )
    )
  );

  registerRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(RegisterActions.registerRequestSuccess),
      tap(() => {
        this.spinnerService.hide();
        this.router.navigate(['/register-success'], { queryParams: { registerState: 'success' }});
      })
    ), { dispatch: false }
  );

  registerRequestFailed$ = createEffect(() =>
    this.actions$.pipe(
      ofType(RegisterActions.registerRequestFailed),
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
