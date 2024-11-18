import { Component, Input, EventEmitter, Output, signal, computed, effect } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Category } from 'src/app/models/category';
import { getValidationMessage } from 'src/app/validations/validations';

@Component({
    selector: 'app-category-form',
    templateUrl: './category-form.component.html',
    styleUrls: ['./category-form.component.scss'],
    standalone: true,
    imports: [FormsModule]
})
export class CategoryFormComponent {
  private _category = signal<Category | null>(null);

  @Input()
  public set category(value: Category | null) {
    this._category.set(value);
    this.categoryName.set(value?.name ?? '');
  }

  public get category(): Category | null {
    return this._category();
  }

  @Input()
  public buttonNames: {
    SubmitButtonText: string;
    CancelButtonText: string;
  } = {
    SubmitButtonText: 'Dodaj',
    CancelButtonText: 'Anuluj'
  };

  @Output()
  public categoryChanged = new EventEmitter<Category>();

  @Output()
  public submitValid = new EventEmitter<void>();

  @Output()
  public cancelClicked = new EventEmitter<void>();

  // Signals for form state
  public categoryName = signal('');
  public categoryNameDirty = signal(false);

  public categoryNameErrors = computed(() => {
    const value = this.categoryName();
    const errors: { [key: string]: any } = {}; // Use `any` to store rich error data
  
    if (!value) {
      errors['required'] = true;
    } else {
      if (value.length < 3) {
        errors['minlength'] = { requiredLength: 3, actualLength: value.length };
      }
      if (value.length > 100) {
        errors['maxlength'] = { requiredLength: 100, actualLength: value.length };
      }
    }
  
    return Object.keys(errors).length > 0 ? errors : null;
  });
  
  constructor() {
    // Emit changes when the category name updates
    effect(() => {
      this.categoryChanged.emit({
        id: this.category?.id ?? '',
        name: this.categoryName(),
        deleted: this.category?.deleted ?? false,
      });
    });
  }

  public getErrorKeys(): string[] {
    return Object.keys(this.categoryNameErrors() ?? {});
  }

  public onSubmit(): void {
    this.categoryNameDirty.set(true);

    if (this.categoryNameErrors()) {
      return; // Form is invalid
    }

    this.submitValid.emit();
  }

  public getErrorMessage(code: { key: string; value: any | null | undefined }): string | null {
    return getValidationMessage(code);
  }

  public cancelClick(): void {
    this.cancelClicked.emit();
  }
}
