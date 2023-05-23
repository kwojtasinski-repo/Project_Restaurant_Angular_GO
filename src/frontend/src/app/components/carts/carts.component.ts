import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { Product } from 'src/app/models/product';
import { fetchCart, finalizeCart, removeProductFromCart } from 'src/app/stores/cart/cart.actions';
import { getCart, getFetchState } from 'src/app/stores/cart/cart.selectors';
import { CartState } from 'src/app/stores/cart/cart.state';

@Component({
  selector: 'app-carts',
  templateUrl: './carts.component.html',
  styleUrls: ['./carts.component.scss']
})
export class CartsComponent implements OnInit {
  public cart$ = this.cartStore.select(getCart);
  public fetchState$ = this.cartStore.select(getFetchState);

  constructor(private cartStore: Store<CartState>) { }
  
  public ngOnInit(): void {
    this.cartStore.dispatch(fetchCart());
  }

  public deleteProduct(product: Product): void {
    this.cartStore.dispatch(removeProductFromCart({ product }));
    this.cartStore.dispatch(fetchCart());
  }

  public calculateTotal(products: Product[] | undefined): number {
    return products ? products.reduce((total, product) => total + product.price, 0) : 0;
  }

  public finalizeOrder(): void {
    // need redirect to url with order
    // think about store will be helpful for loading
    this.cartStore.dispatch(finalizeCart());
  }
}
