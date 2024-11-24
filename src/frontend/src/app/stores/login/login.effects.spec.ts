import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { Observable } from 'rxjs';

import { LoginEffects } from './login.effects';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('LoginEffects', () => {
  let actions$: Observable<any>;
  let effects: LoginEffects;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule],
      providers: [
        LoginEffects,
        provideMockActions(() => actions$)
      ]
    }).compileComponents();

    effects = TestBed.inject(LoginEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
