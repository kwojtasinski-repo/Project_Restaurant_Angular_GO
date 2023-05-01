import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { LoginEffects } from './login.effects';
import { initialState } from './login.reducers';

describe('LoginEffects', () => {
  let actions$: Observable<any>;
  let effects: LoginEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        LoginEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState })
      ]
    });

    effects = TestBed.inject(LoginEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
