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
import { Product } from 'src/app/models/product';
import { ActivatedRoute } from '@angular/router';
import { convertToParamMap } from '@angular/router';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';

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
    productService.get(productInComponent?.id ?? '1').pipe(take(1)).subscribe(p => product = p);
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
});

const fillServiceWithProducts = () => {
  stubbedProducts().forEach(p => productService.add({
    id: p.id,
    name: p.name,
    price: p.price,
    description: p.description ?? '',
    categoryId: p.category?.id ?? '',
  }))
};
