import { AfterViewInit, Component, OnDestroy, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Store } from '@ngrx/store';
import { CategoryState } from 'src/app/stores/category/category.state';
import { Subject, debounceTime, map, takeUntil } from 'rxjs';
import { getError } from 'src/app/stores/category/category.selectors';
import * as CategoryActions from 'src/app/stores/category/category.actions';
import { Category } from 'src/app/models/category';
import { CategoryFormComponent } from '../category-form/category-form.component';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-add-category',
    templateUrl: './add-category.component.html',
    styleUrls: ['./add-category.component.scss'],
    standalone: true,
    imports: [CategoryFormComponent, AsyncPipe]
})
export class AddCategoryComponent implements OnInit, OnDestroy, AfterViewInit {
  public error$ = this.store.select(getError);
  public categoryForm: FormGroup = new FormGroup({});
  private categoryFormValueChanged$ = new Subject();

  constructor(private store: Store<CategoryState>) { }

  public ngOnInit(): void {
    this.categoryForm = new FormGroup({
      categoryName: new FormControl('', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
    });
  }

  public ngAfterViewInit(): void {
    this.categoryForm.valueChanges.pipe(
      debounceTime(10),
      takeUntil(this.categoryFormValueChanged$),
      map((value) => this.store.dispatch(CategoryActions.categoryFormUpdate({
        category: {
          id: '',
          name: value.categoryName,
          deleted: false
        }
      })))
    ).subscribe();
  }

  public onCategoryChange(category: Category): void {
    this.store.dispatch(CategoryActions.categoryFormUpdate({
      category
    }));
  }

  public onSubmit(): void {
    this.store.dispatch(CategoryActions.categoryAddRequestBegin());
  }

  public onCancel(): void {
    this.store.dispatch(CategoryActions.categoryCancelOperation());
  }

  public ngOnDestroy(): void {
    this.categoryFormValueChanged$.unsubscribe();
    this.store.dispatch(CategoryActions.clearErrors());
  }
}
