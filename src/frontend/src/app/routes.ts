import { inject } from "@angular/core";
import { LoginComponent } from "./components/login/login.component";
import { MenuComponent } from "./components/menu/menu.component";
import { ActivatedRouteSnapshot, RouterStateSnapshot, Routes, createUrlTreeFromSnapshot } from "@angular/router";
import { AuthService } from "./services/auth.service";
import { map } from 'rxjs';
import { AddProductsComponent } from "./components/product/add-products/add-products.component";
import { EditProductsComponent } from "./components/product/edit-products/edit-products.component";
import { ViewProductsComponent } from "./components/product/view-products/view-products.component";

const authGuard = (next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot) => {
    const authService = inject(AuthService);

    return authService.isAuthenticated().pipe(
        map((authenticated) => authenticated ? true : createUrlTreeFromSnapshot(next, ['/login']))
    );
};

const adminGuard = (next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot) => {
    const authService = inject(AuthService);

    return authService.getUser().pipe(
        map((user) => user?.role === 'admin' ? true : createUrlTreeFromSnapshot(next, ['/menu']))
    );
};

const authorizedGuard = (next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot) => {
    const authService = inject(AuthService);

    return authService.isAuthenticated().pipe(
        map((authenticated) => authenticated ? createUrlTreeFromSnapshot(next, ['/menu']) : true)
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
            },
            {
                path: 'products/add',
                component: AddProductsComponent,
                canActivate: [adminGuard]
            },
            {
                path: 'products/edit/:id',
                component: EditProductsComponent,
                canActivate: [adminGuard]
            },
            {
                path: 'products/view/:id',
                component: ViewProductsComponent
            }
        ],
    },
    {
        path: 'login',
        component: LoginComponent,
        canActivate: [authorizedGuard]
    },
]

