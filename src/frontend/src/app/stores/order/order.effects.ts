import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { of, catchError, map, mergeMap, tap } from 'rxjs';
import { fetchOrder, fetchOrderFailed, fetchOrderSuccess } from './order.actions';
import { OrderService } from 'src/app/services/order.service';
import { NgxSpinnerService } from 'ngx-spinner';

@Injectable()
export class OrderEffects {
  fectchCart$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchOrder),
      mergeMap((action) => this.orderService.get(action.id)
        .pipe(
          tap(() => this.spinnerService.show()),
          map((order) => fetchOrderSuccess({ order })),
          catchError((err) => of(fetchOrderFailed(err)))
        )
      )
    )
  );

  fetchCartFailed$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchOrderFailed),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  fetchCartSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(fetchOrderSuccess),
        tap(() => this.spinnerService.hide())
    ), {dispatch: false}
  );

  constructor(private actions$: Actions,
    private orderService: OrderService,
    private spinnerService: NgxSpinnerService) {}
}
