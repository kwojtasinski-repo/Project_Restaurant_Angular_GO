import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditProductsComponent } from './edit-products.component';
import { ProductService } from 'src/app/services/product.service';
import { ProductFormComponent } from '../product-form/product-form.component';
import { CategoryService } from 'src/app/services/category.service';
import { fillCategoryServiceWithDefaultValues, InMemoryCategoryService } from 'src/app/unit-test-fixtures/in-memory-category.service';
import { take } from 'rxjs';
import { Category } from 'src/app/models/category';
import { ProductState } from 'src/app/stores/product/product.state';
import { Store } from '@ngrx/store';
import { Product } from 'src/app/models/product';
import { fillProductServiceWithDefaultValues, InMemoryProductService } from 'src/app/unit-test-fixtures/in-memory-product.service';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';
import { createActivatedRouteProvider } from 'src/app/unit-test-fixtures/router-utils';
import { ProductForm } from '../product-form/product-form';

describe('EditProductsComponent', () => {
  let component: EditProductsComponent;
  let fixture: ComponentFixture<EditProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        EditProductsComponent,
        TestSharedModule
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
  let productService: InMemoryProductService;
  let categoryService: InMemoryCategoryService;
  const productId = '1'
  let productForm: ProductForm;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        EditProductsComponent,
        ProductFormComponent,
        TestSharedModule
      ],
      providers: [
        createActivatedRouteProvider({
          id: productId
        })
      ]
    })
    .compileComponents();

    formater = new Intl.NumberFormat('pl-PL', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
    fixture = TestBed.createComponent(EditProductsComponent);
    store = TestBed.inject(Store<ProductState>);
    productService = TestBed.inject(ProductService) as InMemoryProductService;
    categoryService = TestBed.inject(CategoryService) as InMemoryCategoryService;
    fillProductServiceWithDefaultValues(productService);
    fillCategoryServiceWithDefaultValues(categoryService);
    component = fixture.componentInstance;
    fixture.detectChanges();
    productForm = new ProductForm(fixture.nativeElement);
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
    const productName = productForm.getProductNameInput();
    const productDescription = productForm.getProductDescriptionInput();
    const productCost = productForm.getProductCostInput();
    const productCategory = productForm.getProductCategorySelectList();

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
    productForm.getProductNameInput().value = productName;
    productForm.getProductCostInput().value = formater.format(price);
    productForm.getProductDescriptionInput().value = productDescription;
    productForm.getProductCategorySelectList().selectedIndex = 2;

    fixture.detectChanges();

    expect(productForm.getProductNameInput().value).toBe('abc123');
    expect(productForm.getProductCostInput().value).toBe(formater.format(price));
    expect(productForm.getProductDescriptionInput().value).toBe(productDescription);
    expect(categories[productForm.getProductCategorySelectList().selectedIndex].id).toBe(productCategory.id);
  });

  it('should invoke method onSubmit while send form', () => {
    const productName = 'abc123';
    const price = 200;
    const productDescription = 'Description #Product New Value';
    productForm.getProductNameInput().value = productName;
    productForm.getProductCostInput().value = formater.format(price);
    productForm.getProductDescriptionInput().value = productDescription;
    productForm.getProductCategorySelectList().selectedIndex = 2;
    const onSubmit = spyOn(fixture.componentInstance, 'onSubmit').and.callThrough();
    const onDispatch = spyOn(store, 'dispatch').and.callThrough();

    productForm.getProductForm().dispatchEvent(new Event('submit'));

    expect(onSubmit).toHaveBeenCalled();
    expect(onDispatch).toHaveBeenCalled();
  });
});
