import { TestBed } from '@angular/core/testing';

import { CartService } from './cart.service';
import { HttpClientModule } from '@angular/common/http';

describe('CartService', () => {
  let service: CartService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        HttpClientModule
      ],
      providers: [
        {
          provide: "API_URL", useValue: ''
        }
      ]
    });
    service = TestBed.inject(CartService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
