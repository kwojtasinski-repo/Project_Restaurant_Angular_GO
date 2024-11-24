import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { Observable } from 'rxjs';

import { CategoryEffects } from './category.effects';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('CategoryEffects', () => {
  let actions$: Observable<any>;
  let effects: CategoryEffects;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule],
      providers: [
        CategoryEffects,
        provideMockActions(() => actions$)
      ]
    }).compileComponents();

    effects = TestBed.inject(CategoryEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
