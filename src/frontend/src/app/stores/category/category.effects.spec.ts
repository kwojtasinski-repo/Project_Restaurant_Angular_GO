import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { CategoryEffects } from './category.effects';
import { initialState } from './category.reducers';
import { HttpClientModule } from '@angular/common/http';

describe('CategoryEffects', () => {
  let actions$: Observable<any>;
  let effects: CategoryEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        CategoryEffects,
        provideMockActions(() => actions$),
        provideMockStore({ initialState }),
        {
          provide: 'API_URL', useValue: ''
        }
      ],
      imports: [
        HttpClientModule
      ]
    });

    effects = TestBed.inject(CategoryEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
