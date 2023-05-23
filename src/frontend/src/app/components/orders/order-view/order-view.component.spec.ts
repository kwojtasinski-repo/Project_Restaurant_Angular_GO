import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrderViewComponent } from './order-view.component';
import { MoneyPipe } from 'src/app/pipes/money-pipe';
import { RouterTestingModule } from '@angular/router/testing';

describe('OrderViewComponent', () => {
  let component: OrderViewComponent;
  let fixture: ComponentFixture<OrderViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        OrderViewComponent, 
        MoneyPipe 
      ],
      imports: [
        RouterTestingModule
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
