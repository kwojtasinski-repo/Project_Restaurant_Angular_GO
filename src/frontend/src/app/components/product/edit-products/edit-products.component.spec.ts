import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditProductsComponent } from './edit-products.component';
import { RouterTestingModule } from '@angular/router/testing';
import { initialState } from 'src/app/stores/product/product.reducers';
import { provideMockStore, MockStore } from '@ngrx/store/testing';
import { NgxSpinnerModule } from 'ngx-spinner';
import { HttpClientModule } from '@angular/common/http';
import { stubbedCategories, stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';
import productService from 'src/app/unit-test-fixtures/in-memory-product.service';
import { ProductService } from 'src/app/services/product.service';
import { ActivatedRoute, convertToParamMap } from '@angular/router';
import { ProductFormComponent } from '../product-form/product-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { CurrencyFormatterDirective } from 'src/app/directives/currency-formatter-directive';
import { CategoryService } from 'src/app/services/category.service';
import categoryService from 'src/app/unit-test-fixtures/in-memory-category.service';
import { of, take } from 'rxjs';
import { Category } from 'src/app/models/category';
import { ProductState } from 'src/app/stores/product/product.state';
import { MemoizedSelector, Store } from '@ngrx/store';
import { getError } from 'src/app/stores/cart/cart.selectors';

describe('EditProductsComponent', () => {
  let component: EditProductsComponent;
  let fixture: ComponentFixture<EditProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditProductsComponent ],
      imports: [
        RouterTestingModule,
        NgxSpinnerModule,
        HttpClientModule
      ],
      providers: [
        provideMockStore({ initialState }),
        {
          provide: "API_URL", useValue: ''
        },
        {
          provide: ProductService, useValue: productService
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EditProductsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should inform when product is not found', () => {
    const notFoundInformation = fixture.nativeElement.querySelector('.alert.alert-danger > h5');

    expect(notFoundInformation).not.toBeUndefined();
    expect(notFoundInformation).not.toBeNull();
    expect(notFoundInformation.innerHTML).toContain('Produkt nie został znaleziony');
  });
});

describe('EditProductsComponent when product is available', () => {
  let component: EditProductsComponent;
  let fixture: ComponentFixture<EditProductsComponent>;
  let formater: Intl.NumberFormat;
  let store: Store<ProductState>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        EditProductsComponent,
        ProductFormComponent,
        CurrencyFormatterDirective
      ],
      imports: [
        RouterTestingModule,
        NgxSpinnerModule,
        HttpClientModule,
        ReactiveFormsModule
      ],
      providers: [
        provideMockStore({ initialState }),
        {
          provide: ProductService, useValue: productService
        },
        {
          provide: CategoryService, useValue: categoryService
        },
        {
          provide: ActivatedRoute,
          useValue: {
            snapshot: { 
              paramMap:  convertToParamMap({
                id: '1'
              }),
            },
          },
        }
      ]
    })
    .compileComponents();

    formater = new Intl.NumberFormat('pl-PL', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
    fixture = TestBed.createComponent(EditProductsComponent);
    store = TestBed.inject(Store<ProductState>);
    fillProductServiceWithValues();
    fillCategoryServiceWithValues();
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should show product form', () => {
    const form = fixture.nativeElement.querySelector('form');

    expect(form).not.toBeUndefined();
    expect(form).not.toBeNull();
  });

  it('should show product values in inputs', () => {
    let categories: Category[] = [];
    categoryService.getAll().pipe(take(1)).subscribe(c => categories = c);
    const productName = fixture.nativeElement.querySelector('#product-name');
    const productDescription = fixture.nativeElement.querySelector('#product-description');
    const productCost = fixture.nativeElement.querySelector('#product-cost');
    const productCategory = fixture.nativeElement.querySelector('#product-category');

    expect(productName).not.toBeUndefined();
    expect(productName).not.toBeNull();
    expect(productName.value).toEqual(component.product?.name);
    expect(productDescription).not.toBeUndefined();
    expect(productDescription).not.toBeNull();
    expect(productDescription.value).toEqual(component.product?.description);
    expect(productCost).not.toBeUndefined();
    expect(productCost).not.toBeNull();
    expect(productCost.value).toEqual(formater.format(component.product?.price ?? 0));
    expect(productCategory).not.toBeUndefined();
    expect(productCategory).not.toBeNull();
    expect(categories[productCategory.selectedIndex].id - 1).toEqual(component.product?.category?.id ?? 0);
  });

  it('should change inputs while enter new value', () => {
    let categories: Category[] = [];
    categoryService.getAll().pipe(take(1)).subscribe(c => categories = c);
    const productName = 'abc123';
    const price = 200;
    const productDescription = 'Description #Product New Value';
    const productCategory = categories[2];
    fixture.nativeElement.querySelector('#product-name').value = productName;
    fixture.nativeElement.querySelector('#product-cost').value = formater.format(price);
    fixture.nativeElement.querySelector('#product-description').value = productDescription;
    fixture.nativeElement.querySelector('#product-category').selectedIndex = 2;

    fixture.detectChanges();

    expect(fixture.nativeElement.querySelector('#product-name').value).toBe('abc123');
    expect(fixture.nativeElement.querySelector('#product-cost').value).toBe(formater.format(price));
    expect(fixture.nativeElement.querySelector('#product-description').value).toBe(productDescription);
    expect(categories[fixture.nativeElement.querySelector('#product-category').selectedIndex].id).toBe(productCategory.id);
  });

  it('should invoke method onSubmit while send form', () => {
    const productName = 'abc123';
    const price = 200;
    const productDescription = 'Description #Product New Value';
    fixture.nativeElement.querySelector('#product-name').value = productName;
    fixture.nativeElement.querySelector('#product-cost').value = formater.format(price);
    fixture.nativeElement.querySelector('#product-description').value = productDescription;
    fixture.nativeElement.querySelector('#product-category').selectedIndex = 2;
    const onSubmit = spyOn(fixture.componentInstance, 'onSubmit').and.callThrough();
    const onDispatch = spyOn(store, 'dispatch').and.callThrough();

    fixture.nativeElement.querySelector('form').dispatchEvent(new Event('submit'));

    expect(onSubmit).toHaveBeenCalled();
    expect(onDispatch).toHaveBeenCalled();
  });
});

const fillProductServiceWithValues = () => {
  stubbedProducts().forEach(p => productService.add(p))
};

const fillCategoryServiceWithValues = () => {
  stubbedCategories().forEach(c => categoryService.add(c));
}
