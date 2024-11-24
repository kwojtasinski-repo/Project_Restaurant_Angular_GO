import { TestBed } from '@angular/core/testing';

import { CartService } from './cart.service';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';

describe('CartService', () => {
  let service: CartService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
        imports: [TestSharedModule]
    }).compileComponents();
    service = TestBed.inject(CartService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
