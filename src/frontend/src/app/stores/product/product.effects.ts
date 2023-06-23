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
      exhaustMap(([_, product]) => this.productService.add({
        id: '0',
        name: product?.name ?? '',
        price: product?.price ?? 0,
        description: product?.description,
        categoryId: product?.category?.id ?? '0'
      }).pipe(
        map((_) => ProductActions.productAddRequestSuccess()),
        catchError((err) => {
          console.error(err);
          if (err.status === 0) {
            return of(ProductActions.productAddRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
          } else if (err.status >= 500) {
            return of(ProductActions.productAddRequestFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
          }
          return of(ProductActions.productAddRequestFailed({ error: err.error.errors }));
        })
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
      exhaustMap(([_, product]) => this.productService.update({
        id: product?.id ?? '0',
        name: product?.name ?? '',
        price: product?.price ?? 0,
        categoryId: product?.category?.id ?? '0',
        description: product?.description,
      }).pipe(
        map((_) => ProductActions.productAddRequestSuccess()),
        catchError((err) => {
          console.error(err);
          if (err.status === 0) {
            return of(ProductActions.productUpdateRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
          } else if (err.status >= 500) {
            return of(ProductActions.productUpdateRequestFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
          }
          return of(ProductActions.productUpdateRequestFailed({ error: err.error.errors }));
        })
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
