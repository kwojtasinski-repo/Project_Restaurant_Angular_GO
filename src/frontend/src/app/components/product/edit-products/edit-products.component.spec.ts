import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditProductsComponent } from './edit-products.component';
import { initialState } from 'src/app/stores/product/product.reducers';
import { provideMockStore } from '@ngrx/store/testing';
import { NgxSpinnerModule } from 'ngx-spinner';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { stubbedCategories, stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';
import productService from 'src/app/unit-test-fixtures/in-memory-product.service';
import { ProductService } from 'src/app/services/product.service';
import { ActivatedRoute, convertToParamMap, provideRouter } from '@angular/router';
import { ProductFormComponent } from '../product-form/product-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { CurrencyFormatterDirective } from 'src/app/directives/currency-formatter-directive';
import { CategoryService } from 'src/app/services/category.service';
import categoryService from 'src/app/unit-test-fixtures/in-memory-category.service';
import { take } from 'rxjs';
import { Category } from 'src/app/models/category';
import { ProductState } from 'src/app/stores/product/product.state';
import { Store } from '@ngrx/store';
import { Product } from 'src/app/models/product';

describe('EditProductsComponent', () => {
  let component: EditProductsComponent;
  let fixture: ComponentFixture<EditProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
    imports: [NgxSpinnerModule,
        EditProductsComponent],
    providers: [
        provideRouter([]),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        },
        {
            provide: ProductService, useValue: productService
        },
        provideHttpClient(withInterceptorsFromDi())
    ]
})
    .compileComponents();

    fixture = TestBed.createComponent(EditProductsComponent);
    component = fixture.componentInstance;
    component.isLoading = false;
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
  const productId = '1'

  beforeEach(() => {
    TestBed.configureTestingModule({
    imports: [NgxSpinnerModule,
        ReactiveFormsModule,
        EditProductsComponent,
        ProductFormComponent,
        CurrencyFormatterDirective],
    providers: [
        provideRouter([]),
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
                    paramMap: convertToParamMap({
                        id: productId
                    }),
                },
            },
        },
        provideHttpClient(withInterceptorsFromDi())
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
    fixture.detectChanges();
    let categories: Category[] = [];
    let categoriesInComponent: Category[] = [];
    let productInComponent: Product | undefined;
    categoryService.getAll().pipe(take(1)).subscribe(c => categories = c);
    component.categories$.subscribe(c => categoriesInComponent = c);
    component.product$?.subscribe(p => productInComponent = p);
    const productName = fixture.nativeElement.querySelector('#product-name');
    const productDescription = fixture.nativeElement.querySelector('#product-description');
    const productCost = fixture.nativeElement.querySelector('#product-cost');
    const productCategory = fixture.nativeElement.querySelector('#product-category');

    expect(productName).not.toBeUndefined();
    expect(productName).not.toBeNull();
    expect(productName.value).toEqual(productInComponent?.name);
    expect(productDescription).not.toBeUndefined();
    expect(productDescription).not.toBeNull();
    expect(productDescription.value).toEqual(productInComponent?.description);
    expect(productCost).not.toBeUndefined();
    expect(productCost).not.toBeNull();
    expect(productCost.value).toEqual(formater.format(productInComponent?.price ?? 0));
    expect(productCategory).not.toBeUndefined();
    expect(productCategory).not.toBeNull();
    expect((new Number(categories[productCategory.selectedIndex]?.id).valueOf() - 1).toString()).toEqual(productInComponent?.category?.id ?? '0');
    expect(categoriesInComponent.length).toBe(categories.length);
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
  stubbedProducts().forEach(p => productService.add({
    id: p.id,
    name: p.name,
    price: p.price,
    description: p.description,
    categoryId: p.category?.id ?? '1'
  }))
};

const fillCategoryServiceWithValues = () => {
  stubbedCategories().forEach(c => categoryService.add(c));
}
