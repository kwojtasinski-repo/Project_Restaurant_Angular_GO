import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OrderViewComponent } from './order-view.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('OrderViewComponent', () => {
  let component: OrderViewComponent;
  let fixture: ComponentFixture<OrderViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        OrderViewComponent,
        TestSharedModule
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
