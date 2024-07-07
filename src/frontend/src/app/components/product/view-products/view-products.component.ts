import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { finalize, take, tap } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { NgxSpinnerService } from 'ngx-spinner';
import { AuthStateService } from 'src/app/services/auth-state.service';
import { Store } from "@ngrx/store";
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart, clearErrors } from 'src/app/stores/cart/cart.actions';

@Component({
  selector: 'app-view-products',
  templateUrl: './view-products.component.html',
  styleUrls: ['./view-products.component.scss']
})
export class ViewProductsComponent implements OnInit, OnDestroy {
  public product: Product | undefined;
  public isLoading = true;
  public user$ = this.authService.getUser();
  public error: string | undefined;

  constructor(private productService: ProductService, private route: ActivatedRoute, private spinnerService: NgxSpinnerService,
    private authService: AuthStateService, private cartStore: Store<CartState>) { }

  public ngOnInit(): void {
    this.spinnerService.show();
    const id = this.route.snapshot.paramMap.get('id') ?? '';
    this.productService.get(id)
      .pipe(
        take(1),
        tap(() => {
          this.isLoading = true;
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading = false;
          this.spinnerService.hide();
        })
      )
      .subscribe({ next: p => {
          this.product = p
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
        }
      });
  }

  public ngOnDestroy(): void {
    this.cartStore.dispatch(clearErrors());
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
