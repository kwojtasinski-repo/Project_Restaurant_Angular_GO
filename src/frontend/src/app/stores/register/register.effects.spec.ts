import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { Observable } from 'rxjs';

import { RegisterEffects } from './register.effects';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('RegisterEffects', () => {
  let actions$: Observable<any>;
  let effects: RegisterEffects;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule],
      providers: [
        RegisterEffects,
        provideMockActions(() => actions$)
      ]
    }).compileComponents();

    effects = TestBed.inject(RegisterEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
