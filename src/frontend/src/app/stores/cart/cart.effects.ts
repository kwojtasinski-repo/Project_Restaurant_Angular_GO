import { Injectable, inject } from '@angular/core';
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
  private actions$ = inject(Actions);
  private store = inject<Store<CartState>>(Store);
  private cartService = inject(CartService);
  private orderService = inject(OrderService);
  private spinnerService = inject(NgxSpinnerService);
  private router = inject(Router);

  fectchCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CartActions.fetchCart),
      mergeMap(() => this.cartService.getCart()
        .pipe(
          tap(() => this.spinnerService.show()),
          map((cart) => CartActions.fetchCartSuccess({ cart })),
          catchError((err) => {
            console.error(err);
            if (err.status === 0) {
              return of(CartActions.fetchCartFailed({ error: 'Sprawdź połączenie z internetem' }));
            }
            return of(CartActions.fetchCartFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
          })
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
          catchError((err) => {
            console.error(err);
            if (err.status === 0) {
              return of(CartActions.addProductToCartFailed({ error: 'Sprawdź połączenie z internetem' }));
            } else if (err.status >= 500) {
              return of(CartActions.addProductToCartFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
            }
            return of(CartActions.addProductToCartFailed({ error: err.error.errors }));
          })
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
          catchError((err) => {
            console.error(err);
            if (err.status === 0) {
              return of(CartActions.removeProductFromCartFailed({ error: 'Sprawdź połączenie z internetem' }));
            } else if (err.status >= 500) {
              return of(CartActions.removeProductFromCartFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
            }
            return of(CartActions.removeProductFromCartFailed({ error: err.error.errors }));
          })
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
          catchError((err) => {
            console.error(err);
            if (err.status === 0) {
              return of(CartActions.finalizeCartFailed({ error: 'Sprawdź połączenie z internetem' }));
            } else if (err.status >= 500) {
              return of(CartActions.finalizeCartFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
            }

            return of(CartActions.finalizeCartFailed({ error: err.error.errors }));
          })
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
}
