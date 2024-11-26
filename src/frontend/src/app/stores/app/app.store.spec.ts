import { TestBed } from '@angular/core/testing';
import { AppStore } from './app.store';

describe('AppStore', () => {
  let appStore: AppStore;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AppStore],
    });

    appStore = TestBed.inject(AppStore);
  })

  it('should be created', () => {
    expect(appStore).toBeTruthy();
  });
});
