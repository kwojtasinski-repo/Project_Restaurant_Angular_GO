import { Component, OnInit, OnDestroy } from '@angular/core';
import { ProductState } from 'src/app/stores/product/product.state';
import { Store } from "@ngrx/store";
import * as ProductActions from 'src/app/stores/product/product.actions';
import { Product } from 'src/app/models/product';
import { getError } from 'src/app/stores/product/product.selectors';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { BehaviorSubject, EMPTY, Observable, catchError, shareReplay, take } from 'rxjs';

@Component({
  selector: 'app-add-products',
  templateUrl: './add-products.component.html',
  styleUrls: ['./add-products.component.scss']
})
export class AddProductsComponent implements OnInit, OnDestroy {  
  public categories$: Observable<Category[]> = new BehaviorSubject([]);
  public error$ = this.store.select(getError);
  public error = '';

  constructor(private store: Store<ProductState>, private categoryService: CategoryService) { }

  public ngOnInit(): void {
    this.categories$ = this.categoryService.getAll()
      .pipe(
        take(1),
        shareReplay(),
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
  }

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
