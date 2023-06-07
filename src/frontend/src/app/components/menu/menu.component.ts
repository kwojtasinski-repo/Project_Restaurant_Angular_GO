import { Component, OnDestroy, OnInit } from '@angular/core';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { take } from 'rxjs';
import { AuthStateService } from 'src/app/services/auth-state.service';
import { Store } from "@ngrx/store";
import { CartState } from 'src/app/stores/cart/cart.state';
import { addProductToCart, clearErrors } from 'src/app/stores/cart/cart.actions';
import { NgxSpinnerService } from 'ngx-spinner';
import { getError } from 'src/app/stores/cart/cart.selectors';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent implements OnInit, OnDestroy {
  public user$ = this.authService.getUser();
  public products: Product[] = [];
  public productsToShow: Product[] = [];
  public term: string = '';
  public error: string | undefined;
  public cartError$ = this.cartStore.select(getError);

  constructor(private productService: ProductService, private authService: AuthStateService, private cartStore: Store<CartState>,
    private spinnerService: NgxSpinnerService) { }
  
  public ngOnInit(): void {
    this.spinnerService.show();
    this.productService.getAll()
      .pipe(take(1))
      .subscribe({ next: p => {
          this.products = p;
          this.productsToShow = p;
          this.spinnerService.hide();
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          this.spinnerService.hide();
          console.error(error);
        }
      });
  }

  public ngOnDestroy(): void {
    this.cartStore.dispatch(clearErrors());
  }

  public search(term: string): void {
    this.productsToShow = this.products.filter(p => p.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }

  public addToCart(product: Product): void {
    this.cartStore.dispatch(addProductToCart({ product }));
  }
}
