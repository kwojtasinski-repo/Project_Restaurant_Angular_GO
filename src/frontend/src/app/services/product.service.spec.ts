import { TestBed } from '@angular/core/testing';

import { ProductService } from './product.service';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';

describe('ProductService', () => {
  let service: ProductService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule]
    }).compileComponents();
    service = TestBed.inject(ProductService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
