import { TestBed } from '@angular/core/testing';
import { provideMockActions } from '@ngrx/effects/testing';
import { Observable } from 'rxjs';

import { ProductEffects } from './product.effects';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('ProductEffects', () => {
  let actions$: Observable<any>;
  let effects: ProductEffects;

    beforeEach(async () => {
      await TestBed.configureTestingModule({
        imports: [TestSharedModule],
        providers: [
          ProductEffects,
          provideMockActions(() => actions$)
        ]
    }).compileComponents();

    effects = TestBed.inject(ProductEffects);
  });

  it('should be created', () => {
    expect(effects).toBeTruthy();
  });
});
