import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditProductsComponent } from './edit-products.component';
import { RouterTestingModule } from '@angular/router/testing';
import { initialState } from 'src/app/stores/product/product.reducers';
import { provideMockStore } from '@ngrx/store/testing';
import { NgxSpinnerModule } from 'ngx-spinner';

describe('EditProductsComponent', () => {
  let component: EditProductsComponent;
  let fixture: ComponentFixture<EditProductsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EditProductsComponent ],
      imports: [
        RouterTestingModule,
        NgxSpinnerModule
      ],
      providers: [
        provideMockStore({ initialState })
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(EditProductsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
