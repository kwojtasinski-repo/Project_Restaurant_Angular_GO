import { inject } from '@angular/core';
import { Router } from '@angular/router';
import { patchState, signalStore, withMethods, withState } from '@ngrx/signals';
import { catchError, EMPTY, exhaustMap, Observable } from 'rxjs';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { rxMethod } from '@ngrx/signals/rxjs-interop';

type CategoryState = {
  category: Category | null;
  error: string | null;
}

const initialState: CategoryState = {
  category: null,
  error: null
};


export const CategoryStore = signalStore(
  { providedIn: 'root' },
  withState(initialState),
  withMethods((store, router = inject(Router), categoryService = inject(CategoryService)) => {
    const clearErrors = () => {
      patchState(store, { error: null });
    };
    
    const clearCategoryForm = () => {
      patchState(store, { category: null, error: null });
    };

    const handleBackendErrors = (err: any, _: Observable<void>): Observable<void> => {
      if (err.status === 0) {
        patchState(store, { error: 'Sprawdź połączenie z internetem' });
        return EMPTY;
      } else if (err.status >= 500) {
        patchState(store, { error: 'Coś poszło nie tak, spróbuj później' });
        return EMPTY;
      }
      
      patchState(store, { error: err?.error?.errors });
      return EMPTY;
    }

    return {
      addCategory: rxMethod<void> (
        exhaustMap(() => {
          const category = store.category();
          if (!category) {
            patchState(store, { error: 'Kategoria nie może być pusta' });
            return EMPTY;
          }

          clearErrors();
          return categoryService.add(category)
            .pipe(
              catchError((error, caught) => handleBackendErrors(error, caught)),
              exhaustMap(() => {
                router.navigate(['/categories']);
                clearCategoryForm();
                return EMPTY;
              })
            )
        })
      ),
      updateCategory: rxMethod<void> (
        exhaustMap(() => {
          const category = store.category();
          if (!category) {
            patchState(store, { error: 'Kategoria nie może być pusta' });
            return EMPTY;
          }

          clearErrors();
          return categoryService.update(category)
            .pipe(
              catchError((error, caught) => handleBackendErrors(error, caught)),
              exhaustMap(() => {
                router.navigate(['/categories']);
                clearCategoryForm();
                return EMPTY;
              })
            );
        })
      ),
      updateCategoryForm(category: Category): void {
        patchState(store, { category });
      },
      cancelCategoryOperation(): void {
        router.navigate(['/categories']);
        clearCategoryForm();
      },
      clearCategoryForm,
      clearErrors
    }
  })
);
