import { Component, computed, OnDestroy, OnInit, signal, inject } from '@angular/core';
import { take, finalize, catchError, EMPTY, shareReplay } from 'rxjs';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { ActivatedRoute } from '@angular/router';
import { NgxSpinnerService } from 'ngx-spinner';
import { CategoryFormComponent } from '../category-form/category-form.component';
import { CategoryStore } from 'src/app/stores/category/category.store';

@Component({
    selector: 'app-edit-category',
    templateUrl: './edit-category.component.html',
    styleUrls: ['./edit-category.component.scss'],
    standalone: true,
    imports: [CategoryFormComponent]
})
export class EditCategoryComponent implements OnInit, OnDestroy {
  private readonly store = inject(CategoryStore);
  private readonly categoryService = inject(CategoryService);
  private readonly route = inject(ActivatedRoute);
  private readonly spinnerService = inject(NgxSpinnerService);

  public readonly category = signal<Category | null>(null);
  public readonly error = signal<string | null>(null);
  public readonly isLoading = signal<boolean>(true);

  public readonly isError = computed(() => !!this.error() || !!this.storeError());
  public readonly storeError = this.store.error;

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ?? '';

    this.isLoading.set(true);
    this.spinnerService.show();
    this.categoryService.get(id)
      .pipe(
        take(1),
        shareReplay(),
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
    this.store.updateCategoryForm(category);
  }

  public onSubmit(): void {
    this.store.updateCategory();
  }

  public cancelClick(): void {
    this.store.cancelCategoryOperation();
  }

  public ngOnDestroy(): void {
    this.store.clearErrors();
  }
}
