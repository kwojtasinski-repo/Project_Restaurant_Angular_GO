import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { of, catchError, map, mergeMap, tap } from 'rxjs';
import * as OrderActions from './order.actions';
import { OrderService } from 'src/app/services/order.service';
import { NgxSpinnerService } from 'ngx-spinner';

@Injectable()
export class OrderEffects {
  fectchCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(OrderActions.fetchOrder),
      mergeMap((action) => this.orderService.get(action.id)
        .pipe(
          tap(() => this.spinnerService.show()),
          map((order) => OrderActions.fetchOrderSuccess({ order })),
          catchError((err) => {
            console.error(err);
            if (err.status === 0) {
              return of(OrderActions.fetchOrderFailed({ error: 'Sprawdź połączenie z internetem' }));
            }
            return of(OrderActions.fetchOrderFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
          })
        )
      )
    )
  );

  fetchCartFailed$ = createEffect(() =>
    this.actions$.pipe(
      ofType(OrderActions.fetchOrderFailed),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  fetchCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(OrderActions.fetchOrderSuccess),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  constructor(
    private actions$: Actions,
    private orderService: OrderService,
    private spinnerService: NgxSpinnerService
  ) {}
}
