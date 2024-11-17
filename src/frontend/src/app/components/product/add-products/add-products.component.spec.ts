import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddProductsComponent } from './add-products.component';
import { initialState } from 'src/app/stores/product/product.reducers';
import { provideMockStore } from '@ngrx/store/testing';
import { ProductFormComponent } from '../product-form/product-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { CurrencyFormatterDirective } from 'src/app/directives/currency-formatter-directive';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';

describe('AddProductsComponent', () => {
  let component: AddProductsComponent;
  let fixture: ComponentFixture<AddProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
    imports: [ReactiveFormsModule,
        AddProductsComponent, ProductFormComponent, CurrencyFormatterDirective],
    providers: [
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        },
        provideHttpClient(withInterceptorsFromDi()),
    ]
})
    .compileComponents();

    fixture = TestBed.createComponent(AddProductsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
