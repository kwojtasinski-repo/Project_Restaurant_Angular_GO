import { TestBed } from '@angular/core/testing';

import { OrderService } from './order.service';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';

describe('OrderService', () => {
  let service: OrderService;

  beforeEach(() => {
    TestBed.configureTestingModule({
    imports: [],
    providers: [
        {
            provide: 'API_URL', useValue: ''
        },
        provideHttpClient(withInterceptorsFromDi())
    ]
});
    service = TestBed.inject(OrderService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
