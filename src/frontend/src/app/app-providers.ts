import { withInterceptorsFromDi, provideHttpClient } from '@angular/common/http';
import { NgxSpinnerModule } from 'ngx-spinner';
import { RegisterEffects } from './stores/register/register.effects';
import { OrderEffects } from './stores/order/order.effects';
import { CartEffects } from './stores/cart/cart.effects';
import { CategoryEffects } from './stores/category/category.effects';
import { ProductEffects } from './stores/product/product.effects';
import { AppEffects } from './stores/app/app.effects';
import { LoginEffects } from './stores/login/login.effects';
import { EffectsModule } from '@ngrx/effects';
import { registerReducer } from './stores/register/register.reducers';
import { registerStoreName } from './stores/register/register.store.names';
import { orderReducer } from './stores/order/order.reducers';
import { orderStoreName } from './stores/order/order.store.names';
import { cartReducer } from './stores/cart/cart.reducers';
import { cartStoreName } from './stores/cart/cart.store.names';
import { categoryReducer } from './stores/category/category.reducers';
import { categoryStoreName } from './stores/category/category.store.names';
import { productReducer } from './stores/product/product.reducers';
import { productStoreName } from './stores/product/product.store.names';
import { appReducer } from './stores/app/app.reducers';
import { appStoreName } from './stores/app/app.store.names';
import { loginReducer } from './stores/login/login.reducers';
import { loginStoreName } from './stores/login/login.store.names';
import { StoreModule } from '@ngrx/store';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AlertModule } from 'ngx-bootstrap/alert';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { provideAnimations } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { BrowserModule } from '@angular/platform-browser';
import { GlobalErrorHandler } from './global-error-handler';
import { ErrorHandler, importProvidersFrom } from '@angular/core';
import { BsModalService, ModalModule } from 'ngx-bootstrap/modal';
import { API_URL, resolveApiUrl } from './providers/api-url-provider';

export default [
  importProvidersFrom(BrowserModule, AppRoutingModule, CollapseModule.forRoot(), AlertModule.forRoot(), FormsModule, ReactiveFormsModule, StoreModule.forRoot({
    [loginStoreName]: loginReducer,
    [appStoreName]: appReducer,
    [productStoreName]: productReducer,
    [categoryStoreName]: categoryReducer,
    [cartStoreName]: cartReducer,
    [orderStoreName]: orderReducer,
    [registerStoreName]: registerReducer
  }), EffectsModule.forRoot([
    LoginEffects,
    AppEffects,
    ProductEffects,
    CategoryEffects,
    CartEffects,
    OrderEffects,
    RegisterEffects
  ]), NgxSpinnerModule, ModalModule),
  {
    provide: API_URL,
    useFactory: resolveApiUrl,
    multi: true
  },
  BsModalService,
  {
      // processes all errors
      provide: ErrorHandler,
      useClass: GlobalErrorHandler,
  },
  provideAnimations(),
  provideHttpClient(withInterceptorsFromDi())
]
