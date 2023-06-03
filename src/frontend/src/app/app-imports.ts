import { HttpClientModule } from "@angular/common/http";
import { NgxSpinnerModule } from "ngx-spinner";
import { OrderEffects } from "./stores/order/order.effects";
import { CartEffects } from "./stores/cart/cart.effects";
import { CategoryEffects } from "./stores/category/category.effects";
import { ProductEffects } from "./stores/product/product.effects";
import { AppEffects } from "./stores/app/app.effects";
import { LoginEffects } from "./stores/login/login.effects";
import { EffectsModule } from "@ngrx/effects";
import { orderStoreName } from "./stores/order/order.store.names";
import { cartStoreName } from "./stores/cart/cart.store.names";
import { categoryStoreName } from "./stores/category/category.store.names";
import { productStoreName } from "./stores/product/product.store.names";
import { appStoreName } from "./stores/app/app.store.names";
import { loginStoreName } from "./stores/login/login.store.names";
import { loginReducer } from "./stores/login/login.reducers";
import { appReducer } from "./stores/app/app.reducers";
import { orderReducer } from "./stores/order/order.reducers";
import { cartReducer } from "./stores/cart/cart.reducers";
import { categoryReducer } from "./stores/category/category.reducers";
import { productReducer } from "./stores/product/product.reducers";
import { FormsModule, ReactiveFormsModule } from "@angular/forms";
import { StoreModule } from "@ngrx/store";
import { AlertModule } from "ngx-bootstrap/alert";
import { CollapseModule } from "ngx-bootstrap/collapse";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { AppRoutingModule } from "./app-routing.module";
import { BrowserModule } from "@angular/platform-browser";
import { ModalModule } from 'ngx-bootstrap/modal';

export const appImports = [
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
    HttpClientModule,
    ModalModule
]
