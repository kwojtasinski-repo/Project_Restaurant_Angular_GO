import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CategoriesComponent } from './components/category/categories/categories.component';
import { AddCategoryComponent } from './components/category/add-category/add-category.component';
import { ViewCategoryComponent } from './components/category/view-category/view-category.component';
import { ProductsComponent } from './components/product/products/products.component';
import { AddProductsComponent } from './components/product/add-products/add-products.component';
import { ViewProductsComponent } from './components/product/view-products/view-products.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AlertModule } from 'ngx-bootstrap/alert';
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

@NgModule({
  declarations: [
    AppComponent,
    CategoriesComponent,
    AddCategoryComponent,
    ViewCategoryComponent,
    ProductsComponent,
    AddProductsComponent,
    ViewProductsComponent,
    LoginComponent,
    MenuComponent,
    LoginFormDirective,
    HeaderComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    AlertModule.forRoot(),
    FormsModule,
    ReactiveFormsModule,
    StoreModule.forRoot({ [loginStoreName]: loginReducer, [appStoreName]: appReducer }),
    EffectsModule.forRoot([ LoginEffects, AppEffects ])
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
