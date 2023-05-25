import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { OrderEffects } from './order.effects';
import { initialState } from './order.reducers';
import { initialState as initialLoginState } from '../login/login.reducers';
import { getUser } from '../login/login.selectors';

describe('OrderEffects', () => {
  let actions$: Observable<any>;
  let effects: OrderEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        OrderEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        provideMockStore({ initialState: initialLoginState, 
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
        })
      ]
    });

    effects = TestBed.inject(OrderEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
