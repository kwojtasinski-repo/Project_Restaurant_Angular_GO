import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { CartEffects } from './cart.effects';
import { initialState } from './cart.reducers';
import { initialState as initialLoginState } from '../login/login.reducers';
import { getUser } from '../login/login.selectors';
import { HttpClientModule } from '@angular/common/http';

describe('CartEffects', () => {
  let actions$: Observable<any>;
  let effects: CartEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        CartEffects,
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
      ],
      imports: [
        HttpClientModule
      ]
    });

    effects = TestBed.inject(CartEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
