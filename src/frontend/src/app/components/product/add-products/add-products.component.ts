import { Component, OnDestroy } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ProductState } from 'src/app/stores/product/product.state';
import { Store } from "@ngrx/store";
import { productFormClear, productFormUpdate, productRequestBegin } from 'src/app/stores/product/product.actions';
import { getValidationMessage } from 'src/app/validations/validations';
import { Subject, debounceTime, takeUntil } from 'rxjs';

@Component({
  selector: 'app-add-products',
  templateUrl: './add-products.component.html',
  styleUrls: ['./add-products.component.scss']
})
export class AddProductsComponent implements OnDestroy {
  public addProduct: FormGroup;
  public locale: string = 'pl-PL';
  public minimumFractionDigits: number = 2;
  public maximumFractionDigits: number = 2;
  private comma: string = ".";
  private productFormValueChanged$ = new Subject();
  
  constructor(private store: Store<ProductState>) {
    this.addProduct = new FormGroup({
      productName: new FormControl('', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
      productDescription: new FormControl('', Validators.maxLength(5000)),
      productCost: new FormControl('0.00', Validators.compose([Validators.required, Validators.min(0)])),
      productCategory: new FormControl('', Validators.required),
    });
    this.comma = new Intl.NumberFormat(this.locale, { minimumFractionDigits: this.minimumFractionDigits, maximumFractionDigits: this.maximumFractionDigits }).format(0.1).charAt(1);
  }

  public ngAfterViewInit() {
    this.addProduct.valueChanges.pipe(debounceTime(10), takeUntil(this.productFormValueChanged$))
      .subscribe((value) => this.store.dispatch(productFormUpdate({
        product: {
          id: 0,
          name: value.productName,
          price: this.onPriceChange(value.productCost),
          category: {
            id: new Number(value.productCategory).valueOf(),
            name: '',
            deleted: false,
          },
          description: value.productDescription,
          deleted: false
        }
      })));
  }

  public onSubmit() {
    if (this.addProduct.invalid) {
      Object.keys(this.addProduct.controls).forEach(key => {
        this.addProduct.get(key)?.markAsDirty();
      });
      return;
    }
    this.store.dispatch(productRequestBegin());
  }

  public getErrorMessage(error: any): string | null {
    return getValidationMessage(error);
  }

  public ngOnDestroy() {
    this.productFormValueChanged$.complete();
    this.store.dispatch(productFormClear());
  }

  private onPriceChange(value: string): number {
    const newValue = this.comma === "," ? value.replace(this.comma, ".") : value;
    return new Number(newValue.replace(/[^0-9.-]+/g,"")).valueOf();
  }
}
