import { Component, computed, OnDestroy, signal, inject } from '@angular/core';
import { Category } from 'src/app/models/category';
import { CategoryFormComponent } from '../category-form/category-form.component';
import { CategoryStore } from 'src/app/stores/category/category.store';

@Component({
    selector: 'app-add-category',
    templateUrl: './add-category.component.html',
    styleUrls: ['./add-category.component.scss'],
    standalone: true,
    imports: [CategoryFormComponent]
})
export class AddCategoryComponent implements OnDestroy {
  private store = inject(CategoryStore);

  public category = signal<Category | null>(null);
  public isLoading = signal<boolean>(true);

  public isError = computed(() => !!this.storeError());
  public storeError = this.store.error;

  public onCategoryChange(category: Category): void {
    this.store.updateCategoryForm(category);
  }

  public onSubmit(): void {
    this.store.addCategory();
  }

  public cancelClick(): void {
    this.store.cancelCategoryOperation();
  }

  public ngOnDestroy(): void {
    this.store.clearErrors();
  }
}
