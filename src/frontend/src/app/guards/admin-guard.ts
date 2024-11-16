import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, RouterStateSnapshot } from '@angular/router';
import { inject } from '@angular/core';
import { map } from 'rxjs';
import { Store } from '@ngrx/store';
import { LoginState } from '../stores/login/login.state';
import * as LoginSelectors from '../stores/login/login.selectors';

export default (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const loginStore = inject(Store<LoginState>);
    return loginStore.select(LoginSelectors.getUser).pipe(
        map((user) => user?.role === 'admin' ? true : createUrlTreeFromSnapshot(next, ['/menu']))
    );
};
