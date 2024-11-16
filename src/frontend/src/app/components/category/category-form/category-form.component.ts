import { Component, OnDestroy, Input, EventEmitter, Output, ChangeDetectorRef, AfterViewInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Subject, debounceTime, takeUntil } from 'rxjs';
import { Category } from 'src/app/models/category';
import { getValidationMessage } from 'src/app/validations/validations';

@Component({
  selector: 'app-category-form',
  templateUrl: './category-form.component.html',
  styleUrls: ['./category-form.component.scss']
})
export class CategoryFormComponent implements OnDestroy, AfterViewInit {
  @Input()
  public get category(): Category | null | undefined {
    return this._category;
  }

  public set category(category: Category | null | undefined) {
    this._category = category;
    this.assignVariables();
    this.categoryChanged.emit({
        id: this._category?.id ?? '',
        name: this._category?.name ?? '',
        deleted: this._category?.deleted ?? false
      })
  }

  private _category: Category | null | undefined;

  @Input()
  public buttonNames: Array<string> = ['Dodaj', 'Anuluj'];

  @Output()
  public categoryChanged = new EventEmitter<Category>();

  @Output()
  public submitValid = new EventEmitter<any>();

  @Output()
  public cancel = new EventEmitter<any>();

  public categoryForm: FormGroup = new FormGroup({});
  private categoryFormValueChanged$ = new Subject();
  
  constructor(private changeDetector: ChangeDetectorRef) {
    this.assignVariables();
  }

  public ngAfterViewInit(): void {
    this.assignVariables();
    this.changeDetector.detectChanges();
    this.categoryForm.valueChanges.pipe(debounceTime(10), takeUntil(this.categoryFormValueChanged$))
      .subscribe((value) => this.categoryChanged.emit({
          id: this._category?.id ?? '',
          name: value.categoryName,
          deleted: this._category?.deleted ?? false
        })
      );
  }

  public onSubmit(): void {
    if (this.categoryForm.invalid) {
      Object.keys(this.categoryForm.controls).forEach(key => {
        this.categoryForm.get(key)?.markAsDirty();
      });
      return;
    }
    this.submitValid.emit();
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public cancelClick(): void {
    this.cancel.emit();
  }

  public ngOnDestroy(): void {
    this.changeDetector.detach();
    this.categoryFormValueChanged$.complete();
  }

  private assignVariables(): void {
    this.categoryForm = new FormGroup({
      categoryName: new FormControl(this._category?.name ?? '', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
    });
  }
}
