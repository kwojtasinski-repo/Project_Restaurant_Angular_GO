import { TestBed } from '@angular/core/testing';

import { ProductService } from './product.service';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

describe('ProductService', () => {
  let service: ProductService;

  beforeEach(() => {
    TestBed.configureTestingModule({
        imports: [],
        providers: [
            {
                provide: 'API_URL', useValue: ''
            },
            provideHttpClient(withInterceptorsFromDi()),
            provideHttpClientTesting()
        ]
    });
    service = TestBed.inject(ProductService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
