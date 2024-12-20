import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProductFormComponent } from './product-form.component';
import { createProduct } from 'src/app/unit-test-fixtures/products-utils';
import { stubbedCategories } from 'src/app/unit-test-fixtures/categories-utils';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';
import { ProductForm } from './product-form';

describe('ProductFormComponent', () => {
  let component: ProductFormComponent;
  let fixture: ComponentFixture<ProductFormComponent>;
  const categories = stubbedCategories();
  const initialData = {
    product: {
      ...createProduct(0, 'Poduct#1', 100, 'desc', categories[0])
    },
    categories: [
      ...categories
    ]
  }

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ProductFormComponent,
        TestSharedModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ProductFormComponent);
    component = fixture.componentInstance;
    component.product = initialData.product;
    component.categories = initialData.categories;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

describe('ProductFormComponent with init data', () => {
  let formater: Intl.NumberFormat;
  let component: ProductFormComponent;
  let fixture: ComponentFixture<ProductFormComponent>;
  const categories = stubbedCategories();
  const initialData = {
    product: {
      ...createProduct(0, 'Poduct#1', 100, 'desc', categories[0])
    },
    categories: [
      ...categories
    ]
  }
  let productForm: ProductForm; 

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ProductFormComponent,
        TestSharedModule
      ]
    })
    .compileComponents();

    formater = new Intl.NumberFormat('pl-PL', { minimumFractionDigits: 2, maximumFractionDigits: 2 });
    fixture = TestBed.createComponent(ProductFormComponent);
    component = fixture.componentInstance;
    component.product = initialData.product;
    component.categories = initialData.categories;
    fixture.detectChanges();
    productForm = new ProductForm(fixture.nativeElement);
  });

  it('should show form', () => {
    const form = productForm.getProductForm();

    expect(form).not.toBeUndefined();
    expect(form).not.toBeNull();
    expect(form.innerHTML.length).toBeGreaterThan(0);
  });

  it('should set values on form', () => {
    const name = productForm.getProductNameInput();
    const description = productForm.getProductDescriptionInput();
    const cost = productForm.getProductCostInput();
    const category = productForm.getProductCategorySelectList();

    expect(name).not.toBeUndefined();
    expect(name).not.toBeNull();
    expect(name.value).toEqual(component.product?.name);
    expect(description).not.toBeUndefined();
    expect(description).not.toBeNull();
    expect(description.value).toEqual(component.product?.description);
    expect(cost).not.toBeUndefined();
    expect(cost).not.toBeNull();
    expect(cost.value).not.toEqual(component.product?.price);
    expect(category).not.toBeUndefined();
    expect(category).not.toBeNull();
    const categoriesOptions = Array.from(category.querySelectorAll('option'));
    expect(categoriesOptions.length).toEqual(categories.length + 1); // one extra ('Wybierz kategorie')
    expect(category.innerHTML).toContain('Wybierz kategorię');
    expect(categories[category.selectedIndex-1].id).toEqual(component.product?.category?.id ?? '0');
  });

  it('should change values after user action', () => {
    const newProductName = 'newProductName12'
    const newProductDescription = 'newProductDescriptionAbc'
    const newProductCost = 1200;
    const newCategoryIndex = 4;
    
    productForm.changeProductForm(newProductName, newProductDescription, newProductCost, newCategoryIndex + 1);
    fixture.detectChanges();
    
    const name = productForm.getProductNameInput();
    const description = productForm.getProductDescriptionInput();
    const cost = productForm.getProductCostInput();
    const category = productForm.getProductCategorySelectList();
    expect(name).not.toBeUndefined();
    expect(name).not.toBeNull();
    expect(name.value).toEqual(component.productForm?.get('productName')?.value);
    expect(description).not.toBeUndefined();
    expect(description).not.toBeNull();
    expect(description.value).toEqual(component.productForm?.get('productDescription')?.value);
    expect(cost).not.toBeUndefined();
    expect(cost).not.toBeNull();
    expect(cost.value).not.toEqual(formater.format(component.productForm?.get('productCost')?.value));
    expect(category).not.toBeUndefined();
    expect(category).not.toBeNull();
    expect(categories[category.selectedIndex-1].id).toEqual(component.productForm?.get('productCategory')?.value.id ?? 0);
  });

  it('should show errors while left empty values', () => {
    const newProductName = ''
    const newProductDescription = ''
    const newProductCost = '';
    const newCategoryIndex = 0;

    productForm.changeProductForm(newProductName, newProductDescription, newProductCost, newCategoryIndex);
    fixture.detectChanges();    
    
    const errors = productForm.getProductFormErrors();
    expect(errors).not.toBeUndefined();
    expect(errors).not.toBeNull();
    expect(errors.length).toBeGreaterThan(0);
    errors.forEach(e => expect(e.innerHTML).toContain('Pole jest wymagane'));
  });

  it('should show errors while entered invalid value', () => {
    const newProductName = 'a'
    const newProductDescription = 'a'
    const newProductCost = -200;
    const newCategoryIndex = 1;

    productForm.changeProductForm(newProductName, newProductDescription, newProductCost, newCategoryIndex);
    fixture.detectChanges();    
    
    const errors = productForm.getProductFormErrors();
    expect(errors).not.toBeUndefined();
    expect(errors).not.toBeNull();
    expect(errors.length).toBeGreaterThan(0);
    expect(errors.some(e => e.innerHTML.includes('Pole powinno składać się przynajmniej z 3 znaków'))).toBeTrue();
    expect(errors.some(e => e.innerHTML.includes(`Wartość '${newProductCost}' powinna być większa niż '0'`))).toBeTrue();
  });

  it('should show errors while entered too long name and description', () => {
    const newProductName = createLongWord(150);
    const newProductDescription = createLongWord(5500);
    const newProductCost = 200;
    const newCategoryIndex = 1;

    productForm.changeProductForm(newProductName, newProductDescription, newProductCost, newCategoryIndex);
    fixture.detectChanges();    
    
    const errors = productForm.getProductFormErrors();
    expect(errors).not.toBeUndefined();
    expect(errors).not.toBeNull();
    expect(errors.length).toBeGreaterThan(0);
    expect(errors.some(e => e.innerHTML.includes('Pole nie powinno przekroczyć 100 znaków'))).toBeTrue();
    expect(errors.some(e => e.innerHTML.includes('Pole nie powinno przekroczyć 5000 znaków'))).toBeTrue();
  });
});

function createLongWord(size: number) {
  const word: string[] = [];
  for (let i = 0; i < size; i++) {
    word.push('a');
  }
  return word.join('');
}
