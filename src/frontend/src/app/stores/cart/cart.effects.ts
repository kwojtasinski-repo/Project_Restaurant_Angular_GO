import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { of, catchError, exhaustMap, map, mergeMap, tap } from 'rxjs';
import { CartState } from './cart.state';
import { Store } from '@ngrx/store';
import * as CartActions from './cart.actions';
import { CartService } from 'src/app/services/cart.service';
import { OrderService } from 'src/app/services/order.service';
import { NgxSpinnerService } from 'ngx-spinner';
import { Router } from '@angular/router';

@Injectable()
export class CartEffects {
  fectchCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.fetchCart),
      mergeMap(() => this.cartService.getCart()
        .pipe(
          tap(() => this.spinnerService.show()),
          map((cart) => CartActions.fetchCartSuccess({ cart })),
          catchError((err) => of(CartActions.fetchCartFailed(err)))
        )
      )
    )
  );

  fetchCartFailed$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.fetchCartFailed),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  fetchCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.fetchCartSuccess),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  addProductToCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.addProductToCart),
      exhaustMap((action) => this.cartService.add(action.product)
        .pipe(
          map(() => CartActions.addProductToCartSuccess()),
          catchError((err) => of(CartActions.addProductToCartFailed(err)))
        )
      )
    )
  );

  removeProductFromCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.removeProductFromCart),
      exhaustMap((action) => this.cartService.delete(action.cart.id)
        .pipe(
          map(() => CartActions.removeProductFromCartSuccess()),
          catchError((err) => of(CartActions.removeProductFromCartFailed(err)))
        )
      )
    )
  );

  removeProductFromCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.removeProductFromCartSuccess),
      tap(() => this.store.dispatch(CartActions.fetchCart()))
    ), { dispatch: false }
  );

  finalizeCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.finalizeCart),
      exhaustMap(() => this.orderService.addFromCart()
        .pipe(
          map((orderId) => CartActions.finalizeCartSuccess({ orderId })),
          catchError((err) => of(CartActions.finalizeCartFailed(err)))
        )
      )
    )
  );

  finalizeCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.finalizeCartSuccess),
      tap((action) => this.router.navigate(['/orders/view/' + action.orderId]))
    ), { dispatch: false }
  );

  constructor(private actions$: Actions, 
    private store: Store<CartState>, 
    private cartService: CartService, 
    private orderService: OrderService,
    private spinnerService: NgxSpinnerService, 
    private router: Router) {}
}
