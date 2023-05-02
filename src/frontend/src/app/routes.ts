import { inject } from "@angular/core";
import { LoginComponent } from "./components/login/login.component";
import { MenuComponent } from "./components/menu/menu.component";
import { ActivatedRouteSnapshot, RouterStateSnapshot, Routes, createUrlTreeFromSnapshot } from "@angular/router";
import { AuthService } from "./services/auth.service";
import { map } from 'rxjs';

const authGuard = (next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot) => {
    const authService = inject(AuthService);

    return authService.isAuthenticated().pipe(
        map((authenticated) => authenticated ? true : createUrlTreeFromSnapshot(next, ['/login']))
    );
};

export const customRoutes: Routes = [
    {
        path: '',
        canActivate: [authGuard],
        children: [
            {
                path: 'menu',
                component: MenuComponent
            }
        ],
    },
    {
        path: 'login',
        component: LoginComponent
    },
]
