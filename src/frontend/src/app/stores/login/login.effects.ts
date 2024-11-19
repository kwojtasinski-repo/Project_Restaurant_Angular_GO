import { Injectable, inject } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, mergeMap, of, tap, map, withLatestFrom } from 'rxjs';
import { Router } from '@angular/router';
import { Store, select } from '@ngrx/store';
import { LoginState } from './login.state';
import * as LoginActions from './login.actions';
import * as LoginSelectors from './login.selectors';
import { AuthenticationService } from 'src/app/services/authentication.service';
import { NgxSpinnerService } from 'ngx-spinner';

@Injectable()
export class LoginEffects {
  private actions$ = inject(Actions);
  private router = inject(Router);
  private store = inject<Store<LoginState>>(Store);
  private authenticationService = inject(AuthenticationService);
  private spinnerService = inject(NgxSpinnerService);

  reloginRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.reloginRequestSuccess),
      mergeMap(() => of(LoginActions.loginSuccess()))
    )
  );

  loginRequest$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.loginRequest),
      withLatestFrom(this.store.select(LoginSelectors.getCredentials)),
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
      mergeMap(() => {
        this.spinnerService.hide();
        return of(LoginActions.loginSuccess());
      })
    )
  );

  loginSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(LoginActions.loginSuccess),
      withLatestFrom(this.store.pipe(select(LoginSelectors.getLoginPath))),
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
}
