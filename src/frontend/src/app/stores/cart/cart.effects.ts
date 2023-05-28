import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { of, catchError, exhaustMap, map, mergeMap, tap } from 'rxjs';
import { CartState } from './cart.state';
import { Store } from '@ngrx/store';
import { addProductToCart, addProductToCartFailed, addProductToCartSuccess, fetchCart, fetchCartFailed, fetchCartSuccess, finalizeCart, finalizeCartFailed, finalizeCartSuccess, removeProductFromCart, removeProductFromCartFailed, removeProductFromCartSuccess } from './cart.actions';
import { CartService } from 'src/app/services/cart.service';
import { getCart } from './cart.selectors';
import { OrderService } from 'src/app/services/order.service';
import { NgxSpinnerService } from 'ngx-spinner';
import { Router } from '@angular/router';

@Injectable()
export class CartEffects {
  fectchCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchCart),
      mergeMap(() => this.cartService.getCart()
        .pipe(
          tap(() => this.spinnerService.show()),
          map((cart) => fetchCartSuccess({ cart })),
          catchError((err) => of(fetchCartFailed(err)))
        )
      )
    )
  );

  fetchCartFailed$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchCartFailed),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  fetchCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchCartSuccess),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
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
      exhaustMap(([_, cart]) => this.orderService.add(cart)
        .pipe(
          map((orderId) => finalizeCartSuccess({ orderId })),
          catchError((err) => of(finalizeCartFailed(err)))
        )
      )
    )
  );

  finalizeCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(finalizeCartSuccess),
      tap((action) => {
        this.cartService.finalizeCart();
        this.router.navigate(['/orders/view/' + action.orderId]);
      })
    ), { dispatch: false }
  );

  constructor(private actions$: Actions, 
    private store: Store<CartState>, 
    private cartService: CartService, 
    private orderService: OrderService,
    private spinnerService: NgxSpinnerService, 
    private router: Router) {}
}