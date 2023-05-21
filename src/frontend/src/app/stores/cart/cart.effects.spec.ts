import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { CartEffects } from './cart.effects';
import { initialState } from './cart.reducers';

describe('CartEffects', () => {
  let actions$: Observable<any>;
  let effects: CartEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        CartEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState })
      ]
    });

    effects = TestBed.inject(CartEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
