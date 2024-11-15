import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, RouterStateSnapshot } from "@angular/router";
import { AuthStateService } from "../services/auth-state.service";
import { inject } from "@angular/core";
import { map } from "rxjs";

export const adminGuard = (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const authService = inject(AuthStateService);
    return authService.getUser().pipe(
        map((user) => user?.role === 'admin' ? true : createUrlTreeFromSnapshot(next, ['/menu']))
    );
};
