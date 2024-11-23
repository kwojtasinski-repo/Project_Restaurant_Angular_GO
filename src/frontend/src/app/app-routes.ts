import { LoginComponent } from './components/login/login.component';
import { MenuComponent } from './components/menu/menu.component';
import { AddProductsComponent } from './components/product/add-products/add-products.component';
import { EditProductsComponent } from './components/product/edit-products/edit-products.component';
import { ViewProductComponent } from './components/product/view-product/view-product.component';
import { CategoriesComponent } from './components/category/categories/categories.component';
import { AddCategoryComponent } from './components/category/add-category/add-category.component';
import { EditCategoryComponent } from './components/category/edit-category/edit-category.component';
import { CartsComponent } from './components/carts/carts.component';
import { OrderViewComponent } from './components/orders/order-view/order-view.component';
import { MyOrdersComponent } from './components/orders/my-orders/my-orders.component';
import { RegisterComponent } from './components/register/register.component';
import { RegisterSuccessComponent } from './components/register-success/register-success.component';
import authGuard from './guards/auth-guard';
import adminGuard from './guards/admin-guard';
import authorizedGuard from './guards/authorized-guard';
import registerSuccessGuard from './guards/register-success-guard';
import { Routes } from '@angular/router';

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

export default [
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
                component: ViewProductComponent
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
] as Routes
