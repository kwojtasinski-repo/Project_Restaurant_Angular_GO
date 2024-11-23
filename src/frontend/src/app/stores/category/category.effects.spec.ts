import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { CategoryEffects } from './category.effects';
import { initialState } from './category.reducers';
import { provideHttpClient, withInterceptorsFromDi } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';

describe('CategoryEffects', () => {
  let actions$: Observable<any>;
  let effects: CategoryEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
    imports: [],
    providers: [
        CategoryEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        },
        provideHttpClient(withInterceptorsFromDi()),
        provideHttpClientTesting()
    ]
});

    effects = TestBed.inject(CategoryEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
