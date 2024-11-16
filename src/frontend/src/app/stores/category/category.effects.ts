import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import * as CategoryActions from './category.actions';
import { of, catchError, exhaustMap, map } from 'rxjs';
import { Router } from '@angular/router';
import { getCategory } from './category.selectors';
import { CategoryState } from './category.state';
import { Store } from '@ngrx/store';
import { CategoryService } from 'src/app/services/category.service';

@Injectable()
export class CategoryEffects {
  categoryAddRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CategoryActions.categoryAddRequestBegin),
      concatLatestFrom(() => this.store.select(getCategory)),
      exhaustMap(([_, product]) => this.categoryService.add(product!).pipe(
        map((_) => CategoryActions.categoryAddRequestSuccess()),
        catchError((err) => { 
          console.log(err);
          if (err.status === 0) {
            return of(CategoryActions.categoryAddRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
          } else if (err.status >= 500) {
            return of(CategoryActions.categoryAddRequestFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
          }
          return of(CategoryActions.categoryAddRequestFailed({ error: err.error.errors }));
        })
      )),
    )
  );

  categoryAddRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CategoryActions.categoryAddRequestSuccess),
      map(() => {
        this.router.navigate(['/categories']);
        return CategoryActions.categoryFormClear();
      })
    )
  );

  categoryUpdateRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CategoryActions.categoryUpdateRequestBegin),
      concatLatestFrom(() => this.store.select(getCategory)),
      exhaustMap(([_, category]) => this.categoryService.update(category!).pipe(
        map((_) => CategoryActions.categoryAddRequestSuccess()),
        catchError((err) => { 
          console.log(err);
          if (err.status === 0) {
            return of(CategoryActions.categoryUpdateRequestFailed({ error: 'Sprawdź połączenie z internetem' }));
          } else if (err.status >= 500) {
            return of(CategoryActions.categoryUpdateRequestFailed({ error: 'Coś poszło nie tak, spróbuj później' }));
          }
          return of(CategoryActions.categoryUpdateRequestFailed({ error: err.error.errors }));
        })
      )),
    )
  );

  categoryUpdateRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CategoryActions.categoryUpdateRequestSuccess),
      map(() => {
        this.router.navigate(['/categories']);
        return CategoryActions.categoryFormClear();
      })
    )
  );

  categoryCancelOperation$ = createEffect(() =>
    this.actions$.pipe(
      ofType(CategoryActions.categoryCancelOperation),
      map(() => {
        this.router.navigate(['/categories']);
        return CategoryActions.categoryFormClear();
      })
    )
  );

  constructor(
    private actions$: Actions, 
    private store: Store<CategoryState>, 
    private router: Router, 
    private categoryService: CategoryService
  ) {}
}
