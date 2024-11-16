import { Component, OnDestroy, OnInit } from '@angular/core';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { BehaviorSubject, EMPTY, Observable, catchError, finalize, map, shareReplay, take, tap } from 'rxjs';
import { Store } from '@ngrx/store';
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart, clearErrors } from 'src/app/stores/cart/cart.actions';
import { NgxSpinnerService } from 'ngx-spinner';
import { getError } from 'src/app/stores/cart/cart.selectors';
import { LoginState } from 'src/app/stores/login/login.state';
import * as LoginSelectors from 'src/app/stores/login/login.selectors';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent implements OnInit, OnDestroy {
  public user$ = this.loginStore.select(LoginSelectors.getUser);
  public products$: Observable<Product[]> = new BehaviorSubject([]);
  public productsToShow$: Observable<Product[]> = new BehaviorSubject([]);
  public term: string = '';
  public error: string | undefined;
  public cartError$ = this.cartStore.select(getError);

  constructor(private productService: ProductService, private loginStore: Store<LoginState>, private cartStore: Store<CartState>,
    private spinnerService: NgxSpinnerService) { }
  
  public ngOnInit(): void {
    this.products$ = this.productService.getAll()
      .pipe(
        take(1),
        tap(() => this.spinnerService.show()),
        shareReplay(),
        finalize(() => this.spinnerService.hide()),
        catchError((error) => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
          return EMPTY;
        })
      );

    this.productsToShow$ = this.products$;
    this.products$.subscribe();
  }

  public ngOnDestroy(): void {
    this.cartStore.dispatch(clearErrors());
  }

  public search(term: string): void {
    this.productsToShow$ = this.products$.pipe(
      map(products => products.filter(p => p.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase())))
    );
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
