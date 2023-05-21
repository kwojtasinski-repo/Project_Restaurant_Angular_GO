import { Component, OnDestroy, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { NgxSpinnerService } from 'ngx-spinner';
import { Subscription, take } from 'rxjs';
import { Cart } from 'src/app/models/cart';
import { Product } from 'src/app/models/product';
import { fetchCart, finalizeCart, removeProductFromCart } from 'src/app/stores/cart/cart.actions';
import { getCart } from 'src/app/stores/cart/cart.selectors';
import { CartState } from 'src/app/stores/cart/cart.state';

@Component({
  selector: 'app-carts',
  templateUrl: './carts.component.html',
  styleUrls: ['./carts.component.scss']
})
export class CartsComponent implements OnInit, OnDestroy {
  public cart: Cart = { products: [] };
  public isLoading: boolean = true;
  private getCart$: Subscription = new Subscription();
  public cart$ = this.cartStore.select(getCart);

  constructor(private spinnerService: NgxSpinnerService, private cartStore: Store<CartState>) { }
  
  public ngOnInit(): void {
    this.spinnerService.show();
    this.cartStore.dispatch(fetchCart());
    this.getCart$ = this.cartStore.select(getCart)
      .subscribe(c => {
        this.cart = c;
        this.isLoading = false;
        this.spinnerService.hide();
      });
  }

  public ngOnDestroy(): void {
    this.getCart$.unsubscribe();
  }

  public deleteProduct(product: Product): void {
    this.cartStore.dispatch(removeProductFromCart({ product }));
    this.cartStore.dispatch(fetchCart());
  }

  public calculateTotal(): number {
    return this.cart.products.reduce((total, product) => total + product.price, 0);
  }

  public finalizeOrder(): void {
    // need redirect to url with order
    // think about store will be helpful for loading
    this.cartStore.dispatch(finalizeCart());
  }
}
