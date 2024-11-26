import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { provideRouter, RouterLink } from '@angular/router';
import { BsModalService } from 'ngx-bootstrap/modal';
import { NgxSpinnerModule } from 'ngx-spinner';
import { MoneyPipe } from '../pipes/money-pipe';
import { LoginFormDirective } from '../directives/login-form-directive';
import { CurrencyFormatterDirective } from '../directives/currency-formatter-directive';
import { ProductService } from '../services/product.service';
import { InMemoryProductService } from './in-memory-product.service';
import { CategoryService } from '../services/category.service';
import { InMemoryCategoryService } from './in-memory-category.service';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState as registerInitialState } from 'src/app/stores/register/register.reducers';
import { initialState as loginInitialState } from 'src/app/stores/login/login.reducers';
import { initialState as cartInitialState } from 'src/app/stores/cart/cart.reducers';
import { initialState as categoryInitialState } from 'src/app/stores/category/category.reducers';
import { initialState as productInitialState } from 'src/app/stores/product/product.reducers';
import { initialState as orderInitialState } from 'src/app/stores/order/order.reducers';
import { API_URL } from '../providers/api-url-provider';

@NgModule({
    imports: [
        NgxSpinnerModule,
        RouterLink,
        ReactiveFormsModule,
        MoneyPipe,
        LoginFormDirective,
        CurrencyFormatterDirective
    ],
    providers: [
        BsModalService,
        provideRouter([]),
        {
            provide: API_URL, useValue: ''
        },
        {
            provide: ProductService,
            useClass: InMemoryProductService
        },
        {
            provide: CategoryService,
            useClass: InMemoryCategoryService
        },
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting(),
        provideMockStore({ initialState: registerInitialState }),
        provideMockStore({ initialState: cartInitialState }),
        provideMockStore({ initialState: loginInitialState }),
        provideMockStore({ initialState: categoryInitialState }),
        provideMockStore({ initialState: productInitialState }),
        provideMockStore({ initialState: orderInitialState }),
    ]
})
export class TestSharedModule { }
  