import { TestBed } from '@angular/core/testing';

import { OrderService } from './order.service';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';

describe('OrderService', () => {
  let service: OrderService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule]
    }).compileComponents();
    service = TestBed.inject(OrderService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
