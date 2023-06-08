import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Observable, take } from 'rxjs';

import { ViewProductsComponent } from './view-products.component';
import { RouterTestingModule } from '@angular/router/testing';
import { NgxSpinnerModule } from 'ngx-spinner';
import { initialState } from 'src/app/stores/login/login.reducers';
import { provideMockStore } from '@ngrx/store/testing';
import { HttpClientModule } from '@angular/common/http';
import { ProductService } from 'src/app/services/product.service';
import productService from 'src/app/unit-test-fixtures/in-memory-product.service';
import { Category } from 'src/app/models/category';
import { Product } from 'src/app/models/product';
import { ActivatedRoute } from '@angular/router';
import { convertToParamMap } from '@angular/router';
import { MoneyPipe } from 'src/app/pipes/money-pipe';

describe('ViewProductsComponent', () => {
  let component: ViewProductsComponent;
  let fixture: ComponentFixture<ViewProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        ViewProductsComponent,
        MoneyPipe
      ],
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
          provide: ProductService,
          useValue: productService
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

    fixture = TestBed.createComponent(ViewProductsComponent);
    component = fixture.componentInstance;
    spyOn(productService, 'get').and.returnValue(new Observable(o => { o.next(undefined); o.complete(); }));
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should inform when product is not available', () => {
    const notFoundInformation = fixture.nativeElement.querySelector('.alert.alert-danger > h5');
    console.log('notFoundInformation', notFoundInformation)

    expect(notFoundInformation).not.toBeUndefined();
    expect(notFoundInformation).not.toBeNull();
    expect(notFoundInformation.innerHTML.length).toBeGreaterThan(0);
    expect(notFoundInformation.innerHTML).toContain('Produkt nie został znaleziony');
  });
});

describe('ViewProductsComponent when product available', () => {
  let component: ViewProductsComponent;
  let fixture: ComponentFixture<ViewProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        ViewProductsComponent,
        MoneyPipe
      ],
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
          provide: ProductService,
          useValue: productService
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

    fixture = TestBed.createComponent(ViewProductsComponent);
    fillServiceWithProducts();
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should show product when available', () => {
    const productDescription = fixture.nativeElement.querySelector('.product-description');

    expect(fixture.componentInstance).not.toBeNull();
    expect(fixture.componentInstance).not.toBeUndefined();
    expect(fixture.componentInstance.isLoading).not.toBeTrue();
    expect(productDescription.innerHTML).toContain(component.product?.name);
    expect(productDescription.innerHTML).toContain(component.product?.price);
    expect(productDescription.innerHTML).toContain(component.product?.description);
    expect(productDescription.innerHTML).toContain(component.product?.category?.name);
  });

  it('should show infromation when product is deleted', () => {
    let product: Product | undefined;
    productService.get(fixture.componentInstance.product?.id ?? 1).pipe(take(1)).subscribe(p => product = p);
    product!.deleted = true;
    productService.update(product!);
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
});

const fillServiceWithProducts = () => {
  productService.add(createProduct(undefined, 'Produc#1'));
  productService.add(createProduct(undefined, 'Produc#2'));
  productService.add(createProduct(undefined, 'Produc#3'));
  productService.add(createProduct(undefined, 'Produc#4'));
  productService.add(createProduct(undefined, 'Produc#5'));
}

const createProduct = (id: number | undefined = undefined, 
    name: string | undefined = undefined, 
    price: number | undefined = undefined,
    description: number | undefined = undefined,
    category: Category | undefined = undefined) => {
  return { 
    id: id ?? 0,
    name: name ?? 'product',
    category: category ?? {
      id: 1,
      name: 'category',
      deleted: false
    },
    price: price ?? 100,
    description: description ?? 'Desc1234',
    deleted: false
  } as Product
}
