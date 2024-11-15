import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, RouterStateSnapshot } from "@angular/router";
import { AuthStateService } from "../services/auth-state.service";
import { inject } from "@angular/core";
import { map } from "rxjs";

export default (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const authService = inject(AuthStateService);

    return authService.isAuthenticated().pipe(
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
