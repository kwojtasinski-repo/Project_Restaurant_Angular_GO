import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import * as ProductActions from './product.actions';
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
      ofType(ProductActions.productAddRequestBegin),
      concatLatestFrom(() => this.store.select(getProduct)),
      exhaustMap(([_, product]) => this.productService.add(product!).pipe(
        map((_) => ProductActions.productAddRequestSuccess()),
        catchError((err) => of(ProductActions.productAddRequestFailed(err)))
      )),
    )
  );

  productAddRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.productAddRequestSuccess),
      map(() => {
        this.router.navigate(['/menu']);
        return ProductActions.productFormClear();
      })
    )
  );

  productUpdateRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.productUpdateRequestBegin),
      concatLatestFrom(() => this.store.select(getProduct)),
      exhaustMap(([_, product]) => this.productService.update(product!).pipe(
        map((_) => ProductActions.productAddRequestSuccess()),
        catchError((err) => of(ProductActions.productAddRequestFailed(err)))
      )),
    )
  );

  productUpdateRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.productUpdateRequestSuccess),
      map(() => {
        this.router.navigate(['/menu']);
        return ProductActions.productFormClear();
      })
    )
  );

  productCancelOperation$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProductActions.productCancelOperation),
      map(() => {
        this.router.navigate(['/menu']);
        return ProductActions.productFormClear();
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
