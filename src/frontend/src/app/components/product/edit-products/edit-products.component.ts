import { Component, OnDestroy, OnInit, inject } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Store } from '@ngrx/store';
import { BehaviorSubject, EMPTY, Observable, catchError, finalize, forkJoin, map, 
  shareReplay, take, tap } from 'rxjs';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import * as ProductActions from 'src/app/stores/product/product.actions';
import { ProductState } from 'src/app/stores/product/product.state';
import { NgxSpinnerService } from 'ngx-spinner';
import { getError } from 'src/app/stores/product/product.selectors';
import { clearErrors } from 'src/app/stores/cart/cart.actions';
import { CategoryService } from 'src/app/services/category.service';
import { Category } from 'src/app/models/category';
import { ProductFormComponent } from '../product-form/product-form.component';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-edit-products',
    templateUrl: './edit-products.component.html',
    styleUrls: ['./edit-products.component.scss'],
    standalone: true,
    imports: [ProductFormComponent, AsyncPipe]
})
export class EditProductsComponent implements OnInit, OnDestroy {
  private productService = inject(ProductService);
  private route = inject(ActivatedRoute);
  private store = inject<Store<ProductState>>(Store);
  private categoryService = inject(CategoryService);
  private spinnerService = inject(NgxSpinnerService);

  public product$: Observable<Product | undefined> | undefined;
  public categories$: Observable<Category[]> = new BehaviorSubject([]);
  public isLoading = true;
  public error$ = this.store.select(getError);
  public error: string | undefined;

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ?? '';

    const pipGetAll$ = forkJoin([this.productService.get(id), this.categoryService.getAll()])
      .pipe(
        take(1),
        shareReplay(),
        tap(() => {
          this.isLoading = true;
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading = false;
          this.spinnerService.hide();
        }),
        catchError((error) => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
          return EMPTY;
        })
      );

      this.product$ = pipGetAll$.pipe(
        map(([product, _]) => product)
      );
      this.categories$ = pipGetAll$.pipe(
        map(([_, categories]) => categories)
      );
      pipGetAll$.subscribe();
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
