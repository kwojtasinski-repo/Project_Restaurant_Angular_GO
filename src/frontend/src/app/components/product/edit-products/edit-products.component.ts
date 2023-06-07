import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { take } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { productCancelOperation, productFormClear, productFormUpdate, productUpdateRequestBegin } from 'src/app/stores/product/product.actions';
import { ProductState } from 'src/app/stores/product/product.state';
import { NgxSpinnerService } from 'ngx-spinner';
import { getError } from 'src/app/stores/product/product.selectors';

@Component({
  selector: 'app-edit-products',
  templateUrl: './edit-products.component.html',
  styleUrls: ['./edit-products.component.scss']
})
export class EditProductsComponent implements OnInit {
  public product: Product | undefined;
  public isLoading = true;
  public error$ = this.store.select(getError);
  public error: string | undefined;

  constructor(private productService: ProductService, private route: ActivatedRoute, private store: Store<ProductState>, private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    this.spinnerService.show();
    const id = this.route.snapshot.paramMap.get('id') ? new Number(this.route.snapshot.paramMap.get('id')).valueOf() : 0;
    this.productService.get(id)
      .pipe(take(1))
      .subscribe({ next: p => {
          this.product = p;
          if (this.product) {
            this.store.dispatch(productFormUpdate({
              product: this.product
            }));
          }
          this.isLoading = false;
          this.spinnerService.hide();
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          this.spinnerService.hide();
          console.error(error);
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
