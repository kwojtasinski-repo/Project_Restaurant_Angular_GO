import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MenuComponent } from './menu.component';
import { SearchBarComponent } from '../search-bar/search-bar.component';
import { take } from 'rxjs';
import { ProductService } from 'src/app/services/product.service';
import { Product } from 'src/app/models/product';
import { fillProductServiceWithDefaultValues, InMemoryProductService } from 'src/app/unit-test-fixtures/in-memory-product.service';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('MenuComponent', () => {
  let component: MenuComponent;
  let fixture: ComponentFixture<MenuComponent>;
  let productService: InMemoryProductService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        MenuComponent,
        SearchBarComponent,
        TestSharedModule
      ]
    })
    .compileComponents();

    productService = TestBed.inject(ProductService) as InMemoryProductService;
    fixture = TestBed.createComponent(MenuComponent);
    fillProductServiceWithDefaultValues(productService);
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
