import { TestBed } from '@angular/core/testing';
import { provideMockStore } from '@ngrx/store/testing';

import { OrderService } from './order.service';
import { initialState } from '../stores/login/login.reducers';
import { getUser } from '../stores/login/login.selectors';

describe('OrderService', () => {
  let service: OrderService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        provideMockStore({ initialState, 
          selectors: [
            {
              selector: getUser,
              value: {
                id: 1,
                email: 'string',
                role: 'test',
                deleted: null
              }
            }
          ]  
        }),
      ]
    });
    service = TestBed.inject(OrderService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
