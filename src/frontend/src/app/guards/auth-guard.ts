import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, RouterStateSnapshot } from '@angular/router';
import { AuthenticationService } from '../services/authentication.service';
import { Store } from '@ngrx/store';
import { NgxSpinnerService } from 'ngx-spinner';
import * as LoginActions from '../stores/login/login.actions';
import * as LoginSelectors from '../stores/login/login.selectors';
import { LoginState } from '../stores/login/login.state';
import { catchError, exhaustMap, finalize, of } from 'rxjs';
import { User } from '../models/user';
import { HttpErrorResponse } from '@angular/common/http';

export default (next: ActivatedRouteSnapshot, routerStateSnapshot: RouterStateSnapshot) => {
    if (routerStateSnapshot.url === '/') {
        return createUrlTreeFromSnapshot(next, ['/menu']);
    }

    const authenticationService = inject(AuthenticationService);
    const store = inject(Store<LoginState>);
    const spinnerService = inject(NgxSpinnerService);
    store.dispatch(LoginActions.setTargetPath({ path: routerStateSnapshot.url }));
    spinnerService.show();
    return store.select(LoginSelectors.getAuthenticated)
        .pipe(
            exhaustMap((authenticated: boolean) => {
                if (authenticated) {
                    return of(true);
                }

                return authenticationService.getContext()
                    .pipe(
                        exhaustMap((user: User) => {
                            store.dispatch(LoginActions.reloginRequestSuccess({ user }));
                            return of(true);
                        }),
                        catchError((err: HttpErrorResponse) => {
                            if (err.status === 0) {
                                store.dispatch(LoginActions.reloginRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
                            } else if (err.status === 500) {
                                store.dispatch(LoginActions.reloginRequestFailed({ error: 'Coś poszło nie tak, spróbuj ponownie później' }));
                            }
                            return of(createUrlTreeFromSnapshot(next, ['/login']));
                        })
                    );
            }),
            finalize(() => spinnerService.hide())
        )
};
