import { Component, OnDestroy, Input, EventEmitter, Output, ChangeDetectorRef, AfterViewInit, inject } from '@angular/core';
import { FormControl, FormGroup, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Subject, debounceTime, takeUntil } from 'rxjs';
import { Category } from 'src/app/models/category';
import { Product } from 'src/app/models/product';
import { getValidationMessage } from 'src/app/validations/validations';
import { CurrencyFormatterDirective } from '../../../directives/currency-formatter-directive';
import { KeyValuePipe } from '@angular/common';

@Component({
    selector: 'app-product-form',
    templateUrl: './product-form.component.html',
    styleUrls: ['./product-form.component.scss'],
    standalone: true,
    imports: [FormsModule, ReactiveFormsModule, CurrencyFormatterDirective, KeyValuePipe]
})
export class ProductFormComponent implements OnDestroy, AfterViewInit {
  private changeDetector = inject(ChangeDetectorRef);

  @Input()
  public get product(): Product | null | undefined {
    return this._product;
  }

  public set product(product: Product | null | undefined) {
    this._product = product;
    this.assignVariables();
    this.productChanged.emit({
      id: this._product?.id ?? '',
      name: this._product?.name ?? '',
      price: this._product?.price ?? 0,
      category: this._product && this._product.category ? {
        id: this._product.category.id,
        name: this._product.category.name,
        deleted: this._product.category.deleted,
      } : null,
      description: this._product?.description ?? '',
      deleted: this._product?.deleted ?? false
    });
  }

  private _product: Product | null | undefined;

  @Input()
  public categories: Category[] | null = [];

  @Input()
  public buttonNames: Array<string> = ['Dodaj', 'Anuluj'];

  @Output()
  public productChanged = new EventEmitter<Product>();

  @Output()
  public submitValid = new EventEmitter<any>();

  @Output()
  public cancelClicked = new EventEmitter<any>();

  public productForm: FormGroup = new FormGroup({});
  public locale: string = 'pl-PL';
  public minimumFractionDigits: number = 2;
  public maximumFractionDigits: number = 2;
  private localeNumberFormat = new Intl.NumberFormat(this.locale, { minimumFractionDigits: this.minimumFractionDigits, maximumFractionDigits: this.maximumFractionDigits });
  private comma: string = '.';
  private productFormValueChanged$ = new Subject();
  
  constructor() {
    this.assignVariables();
  }

  public ngAfterViewInit(): void {
    this.assignVariables();
    this.changeDetector.detectChanges();
    this.productForm.valueChanges.pipe(debounceTime(10), takeUntil(this.productFormValueChanged$))
      .subscribe((value) => this.productChanged.emit({
          id: this.product?.id ?? '',
          name: value.productName,
          price: this.onPriceChange(value.productCost),
          category: {
            id: value.productCategory.id,
            name: '',
            deleted: false,
          },
          description: value.productDescription,
          deleted: false
        })
      );
  }

  public onSubmit(): void {
    if (this.productForm.invalid) {
      Object.keys(this.productForm.controls).forEach(key => {
        this.productForm.get(key)?.markAsDirty();
      });
      return;
    }
    this.submitValid.emit();
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public cancelClick(): void {
    this.cancelClicked.emit();
  }

  public ngOnDestroy(): void {
    this.changeDetector.detach();
    this.productFormValueChanged$.complete();
  }
  
  public compareCategories(c1: any, c2: any): boolean {
    return c1 && c2 ? c1.id === c2.id : c1 === c2;
  }

  private onPriceChange(value: string): number {
    const newValue = this.comma === ',' ? value.replace(this.comma, '.') : value;
    return new Number(newValue.replace(/[^0-9.-]+/g,'')).valueOf();
  }

  private assignVariables(): void {
    this.productForm = new FormGroup({
      productName: new FormControl(this.product?.name ?? '', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
      productDescription: new FormControl(this.product?.description ?? '', Validators.maxLength(5000)),
      productCost: new FormControl(this.setProductPrice(this.product?.price ?? 0), Validators.compose([Validators.required, Validators.min(0)])),
      productCategory: new FormControl(this.product?.category ?? '', Validators.required),
    });
    this.comma = this.localeNumberFormat.format(0.1).charAt(1);
  }

  private setProductPrice(price: number): string {
    return this.localeNumberFormat.format(price);
  }
}
