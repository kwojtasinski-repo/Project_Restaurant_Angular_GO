import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { EMPTY, Observable, catchError, finalize, shareReplay, take, tap } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { NgxSpinnerService } from 'ngx-spinner';
import { Store } from '@ngrx/store';
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart, clearErrors } from 'src/app/stores/cart/cart.actions';
import { LoginState } from 'src/app/stores/login/login.state';
import * as LoginSelectors from 'src/app/stores/login/login.selectors';
import { MoneyPipe } from '../../../pipes/money-pipe';
import { NgTemplateOutlet, AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-view-products',
    templateUrl: './view-products.component.html',
    styleUrls: ['./view-products.component.scss'],
    standalone: true,
    imports: [NgTemplateOutlet, RouterLink, AsyncPipe, MoneyPipe]
})
export class ViewProductsComponent implements OnInit, OnDestroy {
  public product$: Observable<Product | undefined> | undefined;
  public isLoading = true;
  public user$ = this.loginStore.select(LoginSelectors.getUser);
  public error: string | undefined;

  constructor(private productService: ProductService, private route: ActivatedRoute, private spinnerService: NgxSpinnerService,
    private loginStore: Store<LoginState>, private cartStore: Store<CartState>) { }

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ?? '';
    this.product$ = this.productService.get(id)
      .pipe(
        take(1),
        shareReplay(),
        tap(() => {
          this.isLoading = true;
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading = false;
          this.spinnerService.hide();
        }),
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

    this.product$.subscribe();
  }

  public ngOnDestroy(): void {
    this.cartStore.dispatch(clearErrors());
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
