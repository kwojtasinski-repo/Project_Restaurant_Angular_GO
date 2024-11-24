import { TestBed } from '@angular/core/testing';

import { CategoryService } from './category.service';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';

describe('CategoryService', () => {
  let service: CategoryService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule]
    }).compileComponents();
    service = TestBed.inject(CategoryService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
