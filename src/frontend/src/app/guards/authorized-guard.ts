import { ActivatedRouteSnapshot, createUrlTreeFromSnapshot, RouterStateSnapshot } from "@angular/router";
import { AuthStateService } from "../services/auth-state.service";
import { inject } from "@angular/core";
import { map } from "rxjs";

export default (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const authService = inject(AuthStateService);

    return authService.isAuthenticated().pipe(
        map((authenticated) => authenticated ? createUrlTreeFromSnapshot(next, ['/menu']) : true)
    );
};