import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { of, catchError, exhaustMap, map, mergeMap } from 'rxjs';
import { CartState } from './cart.state';
import { Store } from '@ngrx/store';
import { addProductToCart, addProductToCartFailed, addProductToCartSuccess, fetchCart, fetchCartFailed, fetchCartSuccess, finalizeCart, finalizeCartFailed, finalizeCartSuccess, removeProductFromCart, removeProductFromCartFailed, removeProductFromCartSuccess } from './cart.actions';
import { CartService } from 'src/app/services/cart.service';
import { getCart } from './cart.selectors';
import { OrderService } from 'src/app/services/order.service';

@Injectable()
export class CartEffects {
  fectchCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchCart),
      mergeMap(() => this.cartService.getCart()
        .pipe(
          map((cart) => fetchCartSuccess({ cart })),
          catchError((err) => of(fetchCartFailed(err)))
        )
      )
    )
  );

  addProductToCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(addProductToCart),
      exhaustMap((action) => this.cartService.add(action.product)
        .pipe(
          map(() => addProductToCartSuccess()),
          catchError((err) => of(addProductToCartFailed(err)))
        )
      )
    )
  );

  removeProductFromCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(removeProductFromCart),
      exhaustMap((action) => this.cartService.delete(action.product)
        .pipe(
          map(() => removeProductFromCartSuccess()),
          catchError((err) => of(removeProductFromCartFailed(err)))
        )
      )
    )
  );

  finalizeCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(finalizeCart),
      concatLatestFrom(() => this.store.select(getCart)),
      exhaustMap(([_, cart]) => this.cartService.finalizeCart(cart)
        .pipe(
          map(() => this.orderService.add(cart)),
          map(() => finalizeCartSuccess()),
          catchError((err) => of(finalizeCartFailed(err)))
        )
      )
    )
  );

  constructor(private actions$: Actions, private store: Store<CartState>, private cartService: CartService, private orderService: OrderService) {}
}
