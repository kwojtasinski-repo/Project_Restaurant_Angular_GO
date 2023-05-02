import { inject } from "@angular/core";
import { LoginComponent } from "./components/login/login.component";
import { MenuComponent } from "./components/menu/menu.component";
import { ActivatedRouteSnapshot, Router, RouterStateSnapshot, Routes } from "@angular/router";
import { AuthService } from "./services/auth.service";
import { of } from 'rxjs';

const authGuard = (next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot) => {
    const authService = inject(AuthService);
    const router = inject(Router);

    if (state.url.substring(1) === '') {
        return of(true);
    }

    if (!authService.isAuthenticated()) {
        router.navigate(['']);
        return of(false);
    }

    return of(true);
};

export const customRoutes: Routes = [
    {
        path: '',
        canActivate: [authGuard],
        children: [
            {
                path: '',
                component: LoginComponent,
                data: {
                    ['showHeader']: false
                }
            },
            {
                path: 'menu',
                component: MenuComponent,
                data: {
                    ['showHeader']: true
                }
            }
        ]
    },
]

