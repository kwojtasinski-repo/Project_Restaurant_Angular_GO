import { TestBed } from '@angular/core/testing';

import { CategoryService } from './category.service';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';

describe('CategoryService', () => {
  let service: CategoryService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [TestSharedModule]
    });
    service = TestBed.inject(CategoryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
