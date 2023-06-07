import { Component, OnDestroy } from '@angular/core';
import { ProductState } from 'src/app/stores/product/product.state';
import { Store } from "@ngrx/store";
import * as ProductActions from 'src/app/stores/product/product.actions';
import { Product } from 'src/app/models/product';
import { getError } from 'src/app/stores/product/product.selectors';

@Component({
  selector: 'app-add-products',
  templateUrl: './add-products.component.html',
  styleUrls: ['./add-products.component.scss']
})
export class AddProductsComponent implements OnDestroy {  
  public error$ = this.store.select(getError);

  constructor(private store: Store<ProductState>) { }

  public onProductChange(product: Product): void {
    this.store.dispatch(ProductActions.productFormUpdate({
      product
    }));
  }

  public onSubmit(): void {
    this.store.dispatch(ProductActions.productAddRequestBegin());
  }

  public onCancel(): void {
    this.store.dispatch(ProductActions.productCancelOperation());
  }

  public ngOnDestroy(): void {
    this.store.dispatch(ProductActions.productFormClear());
    this.store.dispatch(ProductActions.clearErrors());
  }
}
