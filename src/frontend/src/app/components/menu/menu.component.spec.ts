import { ComponentFixture, TestBed, fakeAsync } from '@angular/core/testing';

import { MenuComponent } from './menu.component';
import { initialState } from 'src/app/stores/login/login.reducers';
import { initialState as cartInitialState } from 'src/app/stores/cart/cart.reducers';
import { provideMockStore, MockStore } from '@ngrx/store/testing';
import { SearchBarComponent } from '../search-bar/search-bar.component';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { HttpClientModule } from '@angular/common/http';
import { BehaviorSubject, Subject, take } from 'rxjs';
import { User } from 'src/app/models/user';
import { ProductService } from 'src/app/services/product.service';
import productService from 'src/app/unit-test-fixtures/in-memory-product.service';
import { stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';
import { Product } from 'src/app/models/product';

describe('MenuComponent', () => {
  let component: MenuComponent;
  let fixture: ComponentFixture<MenuComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        MenuComponent, 
        SearchBarComponent, 
        MoneyPipe 
      ],
      imports: [
        RouterTestingModule,
        FormsModule,
        HttpClientModule
      ],
      providers: [
        provideMockStore({ initialState }),
        provideMockStore({ initialState: cartInitialState }),
        {
          provide: "API_URL", useValue: ''
        },
        {
          provide: ProductService, useValue: productService
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MenuComponent);
    stubbedProducts().forEach(p => productService.add(p));
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should show products when available', () => {
    let products: Product[] = [];
    productService.getAll().pipe(take(1)).subscribe(p => products = p);

    // curious that I can use here document instead of fixture.nativeElement
    const productsHtml = Array.from(document.querySelectorAll('.d-flex.flex-wrap > div'));

    expect(productsHtml).not.toBeUndefined();
    expect(productsHtml).not.toBeNull();
    expect(productsHtml.length).toEqual(products.length);
    expect(productsHtml.some(p => p.innerHTML.includes(products[1].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[2].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[3].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[4].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[5].name))).toBeTrue();
  });
});
