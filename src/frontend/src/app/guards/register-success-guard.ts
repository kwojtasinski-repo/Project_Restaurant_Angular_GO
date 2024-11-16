import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, RouterStateSnapshot } from "@angular/router";
import { inject } from "@angular/core";
import { map } from "rxjs";
import { Store } from "@ngrx/store";
import { LoginState } from "../stores/login/login.state";
import * as LoginSelectors from "../stores/login/login.selectors";

export default (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const loginStore = inject(Store<LoginState>);

    return loginStore.select(LoginSelectors.getAuthenticated).pipe(
        map((authenticated) => {
            if (authenticated) {
                return createUrlTreeFromSnapshot(next, ['/menu']);
            }

            if (!next.queryParamMap.has('registerState')) {
                return createUrlTreeFromSnapshot(next, ['/register']);
            }

            return true;
        })
    );
};