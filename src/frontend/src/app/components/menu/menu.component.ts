import { Component, OnDestroy, OnInit, WritableSignal, computed, inject, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { EMPTY, catchError, finalize, shareReplay, tap } from 'rxjs';
import { Store } from '@ngrx/store';
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart, clearErrors } from 'src/app/stores/cart/cart.actions';
import { NgxSpinnerService } from 'ngx-spinner';
import { getError } from 'src/app/stores/cart/cart.selectors';
import { LoginState } from 'src/app/stores/login/login.state';
import * as LoginSelectors from 'src/app/stores/login/login.selectors';
import { MoneyPipe } from '../../pipes/money-pipe';
import { SearchBarComponent } from '../search-bar/search-bar.component';
import { RouterLink } from '@angular/router';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-menu',
    templateUrl: './menu.component.html',
    styleUrls: ['./menu.component.scss'],
    standalone: true,
    imports: [RouterLink, SearchBarComponent, AsyncPipe, MoneyPipe]
})
export class MenuComponent implements OnInit, OnDestroy {
  private productService = inject(ProductService);
  private loginStore = inject<Store<LoginState>>(Store);
  private cartStore = inject<Store<CartState>>(Store);
  private spinnerService = inject(NgxSpinnerService);

  public user$ = this.loginStore.select(LoginSelectors.getUser);
  public products: WritableSignal<Product[]> = signal([]);
  public term: WritableSignal<string> = signal('');
  public error: WritableSignal<string | undefined> = signal(undefined);
  public cartError = toSignal(this.cartStore.select(getError));
  
  public productsToShow = computed(() => {
    const term = this.term();
    return this.products().filter(p =>
      p.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase())
    );
  });

  public ngOnInit(): void {
    this.productService.getAll()
      .pipe(
        tap(() => this.spinnerService.show()),
        shareReplay(),
        finalize(() => this.spinnerService.hide()),
        catchError((error) => {
          if (error.status === 0) {
            this.error.set('Sprawdź połączenie z internetem');
          } else if (error.status === 500) {
            this.error.set('Coś poszło nie tak, spróbuj ponownie później');
          }
          console.error(error);
          return EMPTY;
        })
      ).subscribe(products => this.products.set(products));
  }

  public ngOnDestroy(): void {
    this.cartStore.dispatch(clearErrors());
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
