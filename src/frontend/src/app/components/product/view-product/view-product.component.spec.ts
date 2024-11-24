import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Observable, take } from 'rxjs';

import { ViewProductComponent } from './view-product.component';
import { ProductService } from 'src/app/services/product.service';
import { Product } from 'src/app/models/product';
import { stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';
import { InMemoryProductService } from 'src/app/unit-test-fixtures/in-memory-product.service';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';
import { createActivatedRouteProvider } from 'src/app/unit-test-fixtures/router-utils';

describe('ViewProductComponent', () => {
  let component: ViewProductComponent;
  let fixture: ComponentFixture<ViewProductComponent>;
  let productService: InMemoryProductService
  const productId: string = '1';

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ViewProductComponent,
        TestSharedModule
      ],
      providers: [
        createActivatedRouteProvider({
          id: productId
        })
      ]
    }).compileComponents();

    productService = TestBed.inject(ProductService) as InMemoryProductService;
    fixture = TestBed.createComponent(ViewProductComponent);
    component = fixture.componentInstance;
    spyOn(productService, 'get').and.returnValue(new Observable(o => { o.next(undefined); o.complete(); }));
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should inform when product is not available', () => {
    const notFoundInformation = fixture.nativeElement.querySelector('.alert.alert-danger > h5');

    expect(notFoundInformation).not.toBeUndefined();
    expect(notFoundInformation).not.toBeNull();
    expect(notFoundInformation.innerHTML.length).toBeGreaterThan(0);
    expect(notFoundInformation.innerHTML).toContain('Produkt nie został znaleziony');
  });
});

describe('ViewProductsComponent when product available', () => {
  let component: ViewProductComponent;
  let fixture: ComponentFixture<ViewProductComponent>;
  let productService: InMemoryProductService;
  const productId: string = '1';

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ViewProductComponent,
        TestSharedModule
      ],
      providers: [
        createActivatedRouteProvider({
          id: productId
        })
      ]
    }).compileComponents();

    productService = TestBed.inject(ProductService) as InMemoryProductService;
    fixture = TestBed.createComponent(ViewProductComponent);
    fillServiceWithProducts(productService);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should show product when available', () => {
    let productInComponent: Product | undefined;
    component.product$?.subscribe(p => productInComponent = p);
    const productDescription = fixture.nativeElement.querySelector('.product-description');

    expect(fixture.componentInstance).not.toBeNull();
    expect(fixture.componentInstance).not.toBeUndefined();
    expect(fixture.componentInstance.isLoading).not.toBeTrue();
    expect(productDescription.innerHTML).toContain(productInComponent?.name);
    expect(productDescription.innerHTML).toContain(productInComponent?.price);
    expect(productDescription.innerHTML).toContain(productInComponent?.description);
    expect(productDescription.innerHTML).toContain(productInComponent?.category?.name);
  });

  it('should show infromation when product is deleted', () => {
    let productInComponent: Product | undefined;
    component.product$?.subscribe(p => productInComponent = p);
    let product: Product | undefined;
    productService.get(productInComponent?.id ?? productId).pipe(take(1)).subscribe(p => product = p);
    product!.deleted = true;
    productService.update({
      id: product!.id,
      name: product!.name,
      price: product!.price,
      description: product!.description,
      categoryId: product!.category!.id,
    });
    fixture.detectChanges();

    const warningInfo = fixture.nativeElement.querySelector('.text-bg-warning.p-2');

    expect(warningInfo).not.toBeNull();
    expect(warningInfo).not.toBeUndefined();
    expect(warningInfo.innerHTML.length).toBeGreaterThan(0);
    const deleteInfo = warningInfo.querySelector('div');
    expect(deleteInfo).not.toBeNull();
    expect(deleteInfo).not.toBeUndefined();
    expect(deleteInfo.innerHTML.length).toBeGreaterThan(0);
    expect(deleteInfo.innerHTML).toContain('Produkt jest nieużywany');
  });
  
  const fillServiceWithProducts = (productService: ProductService) => {
    stubbedProducts().forEach(p => productService.add({
      id: p.id,
      name: p.name,
      price: p.price,
      description: p.description ?? '',
      categoryId: p.category?.id ?? '',
    }))
  };
});

