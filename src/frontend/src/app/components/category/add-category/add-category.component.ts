import { Component, computed, effect, OnDestroy, signal, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { CategoryState } from 'src/app/stores/category/category.state';
import { getError } from 'src/app/stores/category/category.selectors';
import * as CategoryActions from 'src/app/stores/category/category.actions';
import { Category } from 'src/app/models/category';
import { CategoryFormComponent } from '../category-form/category-form.component';
import { Subject, takeUntil } from 'rxjs';

@Component({
    selector: 'app-add-category',
    templateUrl: './add-category.component.html',
    styleUrls: ['./add-category.component.scss'],
    standalone: true,
    imports: [CategoryFormComponent]
})
export class AddCategoryComponent implements OnDestroy {
  private store = inject<Store<CategoryState>>(Store);
  private destroy$ = new Subject<void>();

  public category = signal<Category | null>(null);
  public isLoading = signal<boolean>(true);

  public isError = computed(() => !!this.storeError());
  public storeError = signal<string | null>(null);

  constructor() {
    effect(() => {
      this.store.select(getError)
        .pipe(takeUntil(this.destroy$))
        .subscribe((err) => { if (err) { this.storeError.set(err)} });
    });
  }

  public onCategoryChange(category: Category): void {
    this.store.dispatch(CategoryActions.categoryFormUpdate({
      category
    }));
  }

  public onSubmit(): void {
    this.store.dispatch(CategoryActions.categoryAddRequestBegin());
  }

  public cancelClick(): void {
    this.store.dispatch(CategoryActions.categoryCancelOperation());
  }

  public ngOnDestroy(): void {
    this.store.dispatch(CategoryActions.clearErrors());
    this.destroy$.next();
    this.destroy$.complete();
  }
}
