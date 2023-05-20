import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { Category } from 'src/app/models/category';
import { CategoryState } from 'src/app/stores/category/category.state';
import { Subject, debounceTime, takeUntil } from 'rxjs';
import { getError } from 'src/app/stores/category/category.selectors';
import { categoryAddRequestBegin, categoryCancelOperation, categoryFormUpdate } from 'src/app/stores/category/category.actions';
import { getValidationMessage } from 'src/app/validations/validations';

@Component({
  selector: 'app-add-category',
  templateUrl: './add-category.component.html',
  styleUrls: ['./add-category.component.scss']
})
export class AddCategoryComponent {
  public error$ = this.store.select(getError);
  public categoryForm: FormGroup = new FormGroup({});
  public category: Category | undefined;
  private categoryFormValueChanged$ = new Subject();

  constructor(private store: Store<CategoryState>) { }

  public ngOnInit(): void {
    this.categoryForm = new FormGroup({
      categoryName: new FormControl('', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
    });
  }

  public ngAfterViewInit(): void {
    this.categoryForm.valueChanges.pipe(debounceTime(10), takeUntil(this.categoryFormValueChanged$))
      .subscribe((value) => this.store.dispatch(categoryFormUpdate({
        category: {
          id: 0,
          name: value.categoryName,
          deleted: false
        }
      })
    ));
  }

  public ngOnDestroy(): void {
    this.categoryFormValueChanged$.unsubscribe();
  }

  public onSubmit(): void {
    if (this.categoryForm.invalid) {
      Object.keys(this.categoryForm.controls).forEach(key => {
        this.categoryForm.get(key)?.markAsDirty();
      });
      return;
    }
    this.store.dispatch(categoryAddRequestBegin());
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public cancelClick(): void {
    this.store.dispatch(categoryCancelOperation());
  }
}
