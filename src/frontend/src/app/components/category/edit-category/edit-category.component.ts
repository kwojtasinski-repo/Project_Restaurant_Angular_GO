import { Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { getError } from 'src/app/stores/category/category.selectors';
import { CategoryState } from 'src/app/stores/category/category.state';
import { getValidationMessage } from 'src/app/validations/validations';
import { debounceTime, takeUntil, Subject, take } from 'rxjs';
import { categoryCancelOperation, categoryFormUpdate, categoryUpdateRequestBegin } from 'src/app/stores/category/category.actions';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { ActivatedRoute } from '@angular/router';
import { NgxSpinnerService } from 'ngx-spinner';

@Component({
  selector: 'app-edit-category',
  templateUrl: './edit-category.component.html',
  styleUrls: ['./edit-category.component.scss']
})
export class EditCategoryComponent implements OnInit, OnDestroy {
  public error$ = this.store.select(getError);
  public categoryForm: FormGroup = new FormGroup({});
  public category: Category | undefined;
  public categoryChanged: boolean = false;
  public isLoading: boolean = true;
  private categoryFormValueChanged$ = new Subject();

  constructor(private store: Store<CategoryState>, private categoryService: CategoryService, private route: ActivatedRoute, private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    this.spinnerService.show();
    this.categoryForm = new FormGroup({
      categoryName: new FormControl('', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
    });
    const id = this.route.snapshot.paramMap.get('id') ? new Number(this.route.snapshot.paramMap.get('id')).valueOf() : 0;
    this.categoryService.get(id)
      .pipe(take(1))
      .subscribe(c => {
        this.category = c;
        if (c) {
          this.categoryForm.get('categoryName')?.setValue(c?.name ?? '');
          this.store.dispatch(categoryFormUpdate({ category: c }));
        }
        this.isLoading = false;
        this.spinnerService.hide();
      });
  }

  public ngAfterViewInit(): void {
    this.categoryForm.valueChanges.pipe(debounceTime(10), takeUntil(this.categoryFormValueChanged$))
      .subscribe((value) => this.store.dispatch(categoryFormUpdate({
        category: {
          id: this.category?.id ?? 0,
          name: value.categoryName,
          deleted: this.category?.deleted ?? false
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
    this.store.dispatch(categoryUpdateRequestBegin());
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public cancelClick(): void {
    this.store.dispatch(categoryCancelOperation());
  }
}
