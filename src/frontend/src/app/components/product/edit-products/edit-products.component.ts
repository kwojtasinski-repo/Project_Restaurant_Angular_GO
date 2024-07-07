import { Component, OnDestroy, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { delay, finalize, forkJoin, take, tap } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import * as ProductActions from 'src/app/stores/product/product.actions';
import { ProductState } from 'src/app/stores/product/product.state';
import { NgxSpinnerService } from 'ngx-spinner';
import { getError } from 'src/app/stores/product/product.selectors';
import { clearErrors } from 'src/app/stores/cart/cart.actions';
import { CategoryService } from 'src/app/services/category.service';
import { Category } from 'src/app/models/category';

@Component({
  selector: 'app-edit-products',
  templateUrl: './edit-products.component.html',
  styleUrls: ['./edit-products.component.scss']
})
export class EditProductsComponent implements OnInit, OnDestroy {
  public product: Product | undefined;
  public isLoading = true;
  public error$ = this.store.select(getError);
  public error: string | undefined;
  public categories: Category[] = [];

  constructor(private productService: ProductService,
    private route: ActivatedRoute,
    private store: Store<ProductState>,
    private categoryService: CategoryService,
    private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ?? '';

    forkJoin([this.productService.get(id), this.categoryService.getAll()])
      .pipe(
        take(1),
        tap(() => {
          this.isLoading = true;
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading = false;
          this.spinnerService.hide();
        })
      )
      .subscribe({ next: ([p, c]) => {
          this.product = p;
          this.categories = c;
          if (this.product) {
            this.store.dispatch(ProductActions.productFormUpdate({
              product: this.product
            }));
          }
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
        }
      });
  }

  public onSubmit() {
    this.store.dispatch(ProductActions.productUpdateRequestBegin());
  }

  public onCancel(): void {
    this.store.dispatch(ProductActions.productCancelOperation());
  }

  public onProductChange(product: Product) {
    this.store.dispatch(ProductActions.productFormUpdate({
      product
    }));
  }

  public ngOnDestroy() {
    this.store.dispatch(ProductActions.productFormClear());
    this.store.dispatch(clearErrors());
  }
}
