import { Component, OnInit, OnDestroy } from '@angular/core';
import { Store } from '@ngrx/store';
import { Cart } from 'src/app/models/cart';
import * as CartActions  from 'src/app/stores/cart/cart.actions';
import { getCart, getFetchState, getFinalizeState } from 'src/app/stores/cart/cart.selectors';
import { CartState } from 'src/app/stores/cart/cart.state';
import { MoneyPipe } from '../../pipes/money-pipe';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { RouterLink } from '@angular/router';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-carts',
    templateUrl: './carts.component.html',
    styleUrls: ['./carts.component.scss'],
    standalone: true,
    imports: [RouterLink, SpinnerButtonComponent, AsyncPipe, MoneyPipe]
})
export class CartsComponent implements OnInit, OnDestroy {
  public carts$ = this.cartStore.select(getCart);
  public fetchState$ = this.cartStore.select(getFetchState);
  public finalizeState$ = this.cartStore.select(getFinalizeState);

  constructor(private cartStore: Store<CartState>) { }
  
  public ngOnInit(): void {
    this.cartStore.dispatch(CartActions.fetchCart());
  }

  public ngOnDestroy(): void {
    this.cartStore.dispatch(CartActions.clearErrors());
  }

  public deleteCart(cart: Cart): void {
    this.cartStore.dispatch(CartActions.removeProductFromCart({ cart }));
  }

  public calculateTotal(carts: Cart[] | null | undefined): number {
    return carts ? carts.reduce((total, cart) => total + new Number(cart.product?.price ?? 0).valueOf(), 0) : 0;
  }

  public finalizeOrder(): void {
    this.cartStore.dispatch(CartActions.finalizeCart());
  }
}
