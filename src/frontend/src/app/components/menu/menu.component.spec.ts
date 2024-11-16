import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MenuComponent } from './menu.component';
import { initialState } from 'src/app/stores/login/login.reducers';
import { initialState as cartInitialState } from 'src/app/stores/cart/cart.reducers';
import { provideMockStore } from '@ngrx/store/testing';
import { SearchBarComponent } from '../search-bar/search-bar.component';
import { FormsModule } from '@angular/forms';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { HttpClientModule } from '@angular/common/http';
import { take } from 'rxjs';
import { ProductService } from 'src/app/services/product.service';
import productService from 'src/app/unit-test-fixtures/in-memory-product.service';
import { stubbedProducts } from 'src/app/unit-test-fixtures/test-utils';
import { Product } from 'src/app/models/product';
import { provideRouter, RouterLink } from '@angular/router';

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
        FormsModule,
        HttpClientModule,
        RouterLink
      ],
      providers: [
        provideRouter([]),
        provideMockStore({ initialState }),
        provideMockStore({ initialState: cartInitialState }),
        {
          provide: 'API_URL', useValue: ''
        },
        {
          provide: ProductService, useValue: productService
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MenuComponent);
    stubbedProducts().forEach(p => productService.add({
      id: p.id,
      name: p.name,
      description: p.description,
      price: p.price,
      categoryId: p.category?.id ?? '0'
    }));
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
    expect(productsHtml.some(p => p.innerHTML.includes(products[0].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[1].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[2].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[3].name))).toBeTrue();
    expect(productsHtml.some(p => p.innerHTML.includes(products[4].name))).toBeTrue();
  });
});
