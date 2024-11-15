import { inject } from "@angular/core";
import { LoginComponent } from "./components/login/login.component";
import { MenuComponent } from "./components/menu/menu.component";
import { ActivatedRouteSnapshot, RouterStateSnapshot, Routes, createUrlTreeFromSnapshot } from "@angular/router";
import { AuthStateService } from "./services/auth-state.service";
import {  map } from 'rxjs';
import { AddProductsComponent } from "./components/product/add-products/add-products.component";
import { EditProductsComponent } from "./components/product/edit-products/edit-products.component";
import { ViewProductsComponent } from "./components/product/view-products/view-products.component";
import { CategoriesComponent } from "./components/category/categories/categories.component";
import { AddCategoryComponent } from "./components/category/add-category/add-category.component";
import { EditCategoryComponent } from "./components/category/edit-category/edit-category.component";
import { CartsComponent } from "./components/carts/carts.component";
import { OrderViewComponent } from "./components/orders/order-view/order-view.component";
import { MyOrdersComponent } from "./components/orders/my-orders/my-orders.component";
import { RegisterComponent } from "./components/register/register.component";
import { RegisterSuccessComponent } from "./components/register-success/register-success.component";
import { authGuard } from "./guards/auth-guard";
import { adminGuard } from "./guards/admin-guard";

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
