import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { ProductEffects } from './product.effects';
import { initialState } from './product.reducers';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

describe('ProductEffects', () => {
  let actions$: Observable<any>;
  let effects: ProductEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
    imports: [],
    providers: [
        ProductEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        },
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting()
    ]
});

    effects = TestBed.inject(ProductEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
