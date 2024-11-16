import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { provideMockStore } from '@ngrx/store/testing';
import { Observable } from 'rxjs';

import { RegisterEffects } from './register.effects';
import { initialState } from './register.reducers';
import { HttpClientModule } from '@angular/common/http';

describe('RegisterEffects', () => {
  let actions$: Observable<any>;
  let effects: RegisterEffects;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        RegisterEffects,
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

    effects = TestBed.inject(RegisterEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
