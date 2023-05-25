import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CategoriesComponent } from './components/category/categories/categories.component';
import { AddCategoryComponent } from './components/category/add-category/add-category.component';
import { AddProductsComponent } from './components/product/add-products/add-products.component';
import { ViewProductsComponent } from './components/product/view-products/view-products.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AlertModule } from 'ngx-bootstrap/alert';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { LoginComponent } from './components/login/login.component';
import { MenuComponent } from './components/menu/menu.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { loginStoreName } from './stores/login/login.store.names';
import { loginReducer } from './stores/login/login.reducers';
import { LoginEffects } from './stores/login/login.effects';
import { LoginFormDirective } from './directives/login-form-directive';
import { HeaderComponent } from './components/header/header.component';
import { appStoreName } from './stores/app/app.store.names';
import { appReducer } from './stores/app/app.reducers';
import { AppEffects } from './stores/app/app.effects';
import { CurrencyFormatterDirective } from './directives/currency-formatter-directive';
import { productStoreName } from './stores/product/product.store.names';
import { productReducer } from './stores/product/product.reducers';
import { ProductEffects } from './stores/product/product.effects';
import { EditProductsComponent } from './components/product/edit-products/edit-products.component';
import { ProductFormComponent } from './components/product/product-form/product-form.component';
import { FooterComponent } from './components/footer/footer.component';
import { NgxSpinnerModule } from 'ngx-spinner';
import { EditCategoryComponent } from './components/category/edit-category/edit-category.component';
import { SearchBarComponent } from './components/search-bar/search-bar.component';
import { categoryStoreName } from './stores/category/category.store.names';
import { categoryReducer } from './stores/category/category.reducers';
import { CategoryEffects } from './stores/category/category.effects';
import { CartEffects } from './stores/cart/cart.effects';
import { cartStoreName } from './stores/cart/cart.store.names';
import { cartReducer } from './stores/cart/cart.reducers';
import { CartsComponent } from './components/carts/carts.component';
import { OrderViewComponent } from './components/orders/order-view/order-view.component';
import { MoneyPipe } from './pipes/money-pipe';
import { MyOrdersComponent } from './components/orders/my-orders/my-orders.component';
import { orderStoreName } from './stores/order/order.store.names';
import { orderReducer } from './stores/order/order.reducers';
import { OrderEffects } from './stores/order/order.effects';
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    CategoriesComponent,
    AddCategoryComponent,
    AddProductsComponent,
    ViewProductsComponent,
    LoginComponent,
    MenuComponent,
    LoginFormDirective,
    HeaderComponent,
    CurrencyFormatterDirective,
    EditProductsComponent,
    ProductFormComponent,
    FooterComponent,
    EditCategoryComponent,
    SearchBarComponent,
    CartsComponent,
    OrderViewComponent,
    MoneyPipe,
    MyOrdersComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    CollapseModule.forRoot(),
    AlertModule.forRoot(),
    FormsModule,
    ReactiveFormsModule,
    StoreModule.forRoot({ 
      [loginStoreName]: loginReducer, 
      [appStoreName]: appReducer, 
      [productStoreName]: productReducer, 
      [categoryStoreName]: categoryReducer, 
      [cartStoreName]: cartReducer, 
      [orderStoreName]: orderReducer
    }),
    EffectsModule.forRoot([ 
      LoginEffects, 
      AppEffects, 
      ProductEffects, 
      CategoryEffects, 
      CartEffects, 
      OrderEffects 
    ]),
    NgxSpinnerModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
