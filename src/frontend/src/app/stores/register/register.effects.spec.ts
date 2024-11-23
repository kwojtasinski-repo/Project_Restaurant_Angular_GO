import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { RegisterEffects } from './register.effects';
import { initialState } from './register.reducers';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

describe('RegisterEffects', () => {
  let actions$: Observable<any>;
  let effects: RegisterEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
    imports: [],
    providers: [
        RegisterEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        },
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting()
    ]
});

    effects = TestBed.inject(RegisterEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
