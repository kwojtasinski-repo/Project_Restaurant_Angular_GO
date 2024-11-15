import { inject } from "@angular/core";
import { LoginComponent } from "./components/login/login.component";
import { MenuComponent } from "./components/menu/menu.component";
import { ActivatedRouteSnapshot, RouterStateSnapshot, Routes, createUrlTreeFromSnapshot } from "@angular/router";
import { AuthStateService } from "./services/auth-state.service";
import { catchError, concatMap, exhaustMap, finalize, map, of } from 'rxjs';
import { AddProductsComponent } from "./components/product/add-products/add-products.component";
import { EditProductsComponent } from "./components/product/edit-products/edit-products.component";
import { ViewProductsComponent } from "./components/product/view-products/view-products.component";
import { CategoriesComponent } from "./components/category/categories/categories.component";
import { AddCategoryComponent } from "./components/category/add-category/add-category.component";
import { EditCategoryComponent } from "./components/category/edit-category/edit-category.component";
import { CartsComponent } from "./components/carts/carts.component";
import { OrderViewComponent } from "./components/orders/order-view/order-view.component";
import { MyOrdersComponent } from "./components/orders/my-orders/my-orders.component";
import { LoginState } from "./stores/login/login.state";
import { Store } from '@ngrx/store';
import * as LoginActions from "./stores/login/login.actions";
import { RegisterComponent } from "./components/register/register.component";
import { RegisterSuccessComponent } from "./components/register-success/register-success.component";
import { AuthenticationService } from "./services/authentication.service";
import { User } from "./models/user";
import { HttpErrorResponse } from "@angular/common/http";
import { NgxSpinnerService } from "ngx-spinner";

const authGuard = (next: ActivatedRouteSnapshot, routerStateSnapshot: RouterStateSnapshot) => {
    if (routerStateSnapshot.url === '/') {
        return createUrlTreeFromSnapshot(next, ['/menu']);
    }

    const authStateService = inject(AuthStateService);
    const authenticationService = inject(AuthenticationService);
    const store = inject(Store<LoginState>);
    const spinnerService = inject(NgxSpinnerService);
    store.dispatch(LoginActions.setTargetPath({ path: routerStateSnapshot.url }));
    spinnerService.show();
    return authStateService.isAuthenticated()
        .pipe(
            concatMap((authenticated: boolean) => {
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

const adminGuard = (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const authService = inject(AuthStateService);
    return authService.getUser().pipe(
        map((user) => user?.role === 'admin' ? true : createUrlTreeFromSnapshot(next, ['/menu']))
    );
};

const authorizedGuard = (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
    const authService = inject(AuthStateService);

    return authService.isAuthenticated().pipe(
        map((authenticated) => authenticated ? createUrlTreeFromSnapshot(next, ['/menu']) : true)
    );
};

const registerSuccessGuard = (next: ActivatedRouteSnapshot, _: RouterStateSnapshot) => {
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

const adminRoutes = [
    {
        path: '',
        canActivate: [adminGuard],
        children: [
            {
                path: 'products/add',
                component: AddProductsComponent
            },
            {
                path: 'products/edit/:id',
                component: EditProductsComponent
            },
            {
                path: 'categories',
                component: CategoriesComponent
            },
            {
                path: 'categories/add',
                component: AddCategoryComponent
            },
            {
                path: 'categories/edit/:id',
                component: EditCategoryComponent
            }
        ]
    }
];

export const appRoutes: Routes = [
    {
        path: '',
        canActivate: [authGuard],
        children: [
            {
                path: 'menu',
                component: MenuComponent
            },
            {
                path: 'products/view/:id',
                component: ViewProductsComponent
            },
            {
                path: 'cart',
                component: CartsComponent
            },
            {
                path: 'orders/view/:id',
                component: OrderViewComponent
            },
            {
                path: 'orders/my',
                component: MyOrdersComponent
            },
            ...adminRoutes,
        ]
    },
    {
        path: 'login',
        component: LoginComponent,
        canActivate: [authorizedGuard],
        data: { hideNavBar: true }
    },
    {
        path: 'register',
        component: RegisterComponent,
        canActivate: [authorizedGuard],
        data: { hideNavBar: true }
    },
    {
        path: 'register-success',
        component: RegisterSuccessComponent,
        canActivate: [registerSuccessGuard],
        data: { registerState: 'success', hideNavBar: true }
    },
    {
        path: '**',
        redirectTo: 'menu' // think if page not found is needed
    },
]
