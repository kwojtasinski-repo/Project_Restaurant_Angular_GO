import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { productFormClear, productRequestBegin, productRequestFailed, productRequestSuccess } from './product.actions';
import { of, catchError, exhaustMap, map } from 'rxjs';
import { Router } from '@angular/router';
import { ProductService } from 'src/app/services/product.service';
import { getProduct } from './product.selectors';
import { ProductState } from './product.state';
import { Store } from '@ngrx/store';

@Injectable()
export class ProductEffects {
  productRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productRequestBegin),
      concatLatestFrom(() => this.store.select(getProduct)),
      exhaustMap(([_, product]) => this.productService.add(product!).pipe(
        map((_) => productRequestSuccess()),
        catchError((err) => of(productRequestFailed(err)))
      )),
    )
  );

  productRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productRequestSuccess),
      map(() => {
        this.router.navigate(['/menu']);
        return productFormClear();
      })
    )
  );

  constructor(private actions$: Actions, private store: Store<ProductState>, private router: Router, private productService: ProductService) {}
}
