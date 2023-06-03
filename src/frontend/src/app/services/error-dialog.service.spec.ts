import { TestBed } from '@angular/core/testing';

import { ErrorDialogService } from './error-dialog.service';
import { BsModalService } from 'ngx-bootstrap/modal';

describe('ErrorDialogService', () => {
  let service: ErrorDialogService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        BsModalService
      ]
    });
    service = TestBed.inject(ErrorDialogService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
