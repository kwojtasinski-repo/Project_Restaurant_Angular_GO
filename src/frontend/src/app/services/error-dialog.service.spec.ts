import { TestBed } from '@angular/core/testing';

import { ErrorDialogService } from './error-dialog.service';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';

describe('ErrorDialogService', () => {
  let service: ErrorDialogService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        TestSharedModule
      ]
    }).compileComponents();
    service = TestBed.inject(ErrorDialogService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
