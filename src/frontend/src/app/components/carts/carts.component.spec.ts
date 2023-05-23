import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CartsComponent } from './carts.component';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from 'src/app/stores/cart/cart.reducers';
import { NgxSpinnerModule } from 'ngx-spinner';

describe('CartsComponent', () => {
  let component: CartsComponent;
  let fixture: ComponentFixture<CartsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        CartsComponent 
      ],
      imports: [
        NgxSpinnerModule
      ],
      providers: [
        provideMockStore({ initialState })
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CartsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
