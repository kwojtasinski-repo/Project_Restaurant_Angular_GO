import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { LoginEffects } from './login.effects';
import { initialState } from './login.reducers';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';

describe('LoginEffects', () => {
  let actions$: Observable<any>;
  let effects: LoginEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
    imports: [],
    providers: [
        LoginEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        },
        provideHttpClient(withInterceptorsFromDi())
    ]
});

    effects = TestBed.inject(LoginEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
