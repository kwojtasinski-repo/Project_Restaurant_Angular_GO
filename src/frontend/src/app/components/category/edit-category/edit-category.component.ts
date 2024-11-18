import { Component, computed, effect, OnDestroy, OnInit, signal } from '@angular/core';
import { Store } from '@ngrx/store';
import { getError } from 'src/app/stores/category/category.selectors';
import { CategoryState } from 'src/app/stores/category/category.state';
import { take, tap, finalize, catchError, EMPTY, shareReplay } from 'rxjs';
import * as CategoryActions from 'src/app/stores/category/category.actions';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { ActivatedRoute } from '@angular/router';
import { NgxSpinnerService } from 'ngx-spinner';
import { CategoryFormComponent } from '../category-form/category-form.component';

@Component({
    selector: 'app-edit-category',
    templateUrl: './edit-category.component.html',
    styleUrls: ['./edit-category.component.scss'],
    standalone: true,
    imports: [CategoryFormComponent]
})
export class EditCategoryComponent implements OnInit, OnDestroy {
  public category = signal<Category | null>(null);
  public error = signal<string | null>(null);
  public isLoading = signal<boolean>(true);

  public isError = computed(() => !!this.error() || !!this.storeError());
  public storeError = signal<string | null>(null);

  constructor(private store: Store<CategoryState>, private categoryService: CategoryService, private route: ActivatedRoute, 
    private spinnerService: NgxSpinnerService) {
      effect(() => {
        this.store.select(getError).subscribe((err) => { if (err) { this.storeError.set(err)} });
      });
    }

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ?? '';

    this.categoryService.get(id)
      .pipe(
        take(1),
        shareReplay(),
        tap(() => {
          this.isLoading.set(true);
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading.set(false);
          this.spinnerService.hide();
        }),
        catchError((error) => {
          if (error.status === 0) {
            this.error.set('Sprawdź połączenie z internetem');
          } else if (error.status === 500) {
            this.error.set('Coś poszło nie tak, spróbuj ponownie później');
          }
          console.error(error);
          return EMPTY;
        })
      ).subscribe(category => {
        this.category.set(category ?? null);
      });
  }

  public onCategoryChange(category: Category): void {
    this.store.dispatch(CategoryActions.categoryFormUpdate({
      category
    }));
  }

  public onSubmit(): void {
    this.store.dispatch(CategoryActions.categoryUpdateRequestBegin());
  }

  public cancelClick(): void {
    this.store.dispatch(CategoryActions.categoryCancelOperation());
  }

  public ngOnDestroy(): void {
    this.store.dispatch(CategoryActions.clearErrors());
  }
}
