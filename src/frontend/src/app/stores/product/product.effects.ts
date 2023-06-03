import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { productFormClear, productAddRequestBegin, productAddRequestFailed, productAddRequestSuccess, productUpdateRequestBegin, productUpdateRequestSuccess, productCancelOperation } from './product.actions';
import { of, catchError, exhaustMap, map } from 'rxjs';
import { Router } from '@angular/router';
import { ProductService } from 'src/app/services/product.service';
import { getProduct } from './product.selectors';
import { ProductState } from './product.state';
import { Store } from '@ngrx/store';

@Injectable()
export class ProductEffects {
  productAddRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productAddRequestBegin),
      concatLatestFrom(() => this.store.select(getProduct)),
      exhaustMap(([_, product]) => this.productService.add(product!).pipe(
        map((_) => productAddRequestSuccess()),
        catchError((err) => of(productAddRequestFailed(err)))
      )),
    )
  );

  productAddRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productAddRequestSuccess),
      map(() => {
        this.router.navigate(['/menu']);
        return productFormClear();
      })
    )
  );

  productUpdateRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productUpdateRequestBegin),
      concatLatestFrom(() => this.store.select(getProduct)),
      exhaustMap(([_, product]) => this.productService.update(product!).pipe(
        map((_) => productAddRequestSuccess()),
        catchError((err) => of(productAddRequestFailed(err)))
      )),
    )
  );

  productUpdateRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productUpdateRequestSuccess),
      map(() => {
        this.router.navigate(['/menu']);
        return productFormClear();
      })
    )
  );

  productCancelOperation$ = createEffect(() =>
    this.actions$.pipe(
      ofType(productCancelOperation),
      map(() => {
        this.router.navigate(['/menu']);
        return productFormClear();
      })
    )
  );

  constructor(
    private actions$: Actions, 
    private store: Store<ProductState>, 
    private router: Router, 
    private productService: ProductService
  ) {}
}
