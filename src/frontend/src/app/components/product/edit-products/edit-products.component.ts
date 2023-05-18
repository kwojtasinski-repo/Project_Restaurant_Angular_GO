import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { take } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { productCancelOperation, productFormClear, productFormUpdate, productUpdateRequestBegin } from 'src/app/stores/product/product.actions';
import { ProductState } from 'src/app/stores/product/product.state';

@Component({
  selector: 'app-edit-products',
  templateUrl: './edit-products.component.html',
  styleUrls: ['./edit-products.component.scss']
})
export class EditProductsComponent implements OnInit {
  public product: Product | undefined;

  constructor(private productService: ProductService, private route: ActivatedRoute, private store: Store<ProductState>) { }

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ? new Number(this.route.snapshot.paramMap.get('id')).valueOf() : 0;
    this.productService.get(id)
      .pipe(take(1))
      .subscribe(p => {
        this.product = p;
        if (this.product) {
          this.store.dispatch(productFormUpdate({
            product: this.product
          }));
        }
      });
  }

  public onSubmit() {
    this.store.dispatch(productUpdateRequestBegin());
  }

  public onCancel(): void {
    this.store.dispatch(productCancelOperation());
  }

  public onProductChange(product: Product) {
    this.store.dispatch(productFormUpdate({
      product
    }));
  }

  public ngOnDestroy() {
    this.store.dispatch(productFormClear());
  }
}
