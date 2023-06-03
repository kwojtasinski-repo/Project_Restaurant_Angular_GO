import { Injectable } from '@angular/core';
import { Actions, concatLatestFrom, createEffect, ofType } from '@ngrx/effects';
import { categoryFormClear, categoryAddRequestBegin, categoryAddRequestSuccess, categoryAddRequestFailed, categoryUpdateRequestSuccess, categoryUpdateRequestBegin, categoryCancelOperation } from './category.actions';
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
      ofType(categoryAddRequestBegin),
      concatLatestFrom(() => this.store.select(getCategory)),
      exhaustMap(([_, product]) => this.categoryService.add(product!).pipe(
        map((_) => categoryAddRequestSuccess()),
        catchError((err) => of(categoryAddRequestFailed(err)))
      )),
    )
  );

  categoryAddRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(categoryAddRequestSuccess),
      map(() => {
        this.router.navigate(['/menu']);
        return categoryFormClear();
      })
    )
  );

  categoryUpdateRequestBegin$ = createEffect(() =>
    this.actions$.pipe(
      ofType(categoryUpdateRequestBegin),
      concatLatestFrom(() => this.store.select(getCategory)),
      exhaustMap(([_, category]) => this.categoryService.update(category!).pipe(
        map((_) => categoryAddRequestSuccess()),
        catchError((err) => of(categoryAddRequestFailed(err)))
      )),
    )
  );

  categoryUpdateRequestSuccess$ = createEffect(() =>
    this.actions$.pipe(
      ofType(categoryUpdateRequestSuccess),
      map(() => {
        this.router.navigate(['/categories']);
        return categoryFormClear();
      })
    )
  );

  categoryCancelOperation$ = createEffect(() =>
    this.actions$.pipe(
      ofType(categoryCancelOperation),
      map(() => {
        this.router.navigate(['/categories']);
        return categoryFormClear();
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
