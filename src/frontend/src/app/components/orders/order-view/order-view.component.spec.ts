import { ComponentFixture, TestBed } from '@angular/core/testing';
import { provideMockStore } from '@ngrx/store/testing';

import { OrderViewComponent } from './order-view.component';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { initialState } from 'src/app/stores/order/order.reducers';
import { provideRouter } from '@angular/router';

describe('OrderViewComponent', () => {
  let component: OrderViewComponent;
  let fixture: ComponentFixture<OrderViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
    imports: [
        OrderViewComponent,
        MoneyPipe
    ],
    providers: [
        provideRouter([]),
        provideMockStore({ initialState }),
    ]
})
    .compileComponents();

    fixture = TestBed.createComponent(OrderViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
