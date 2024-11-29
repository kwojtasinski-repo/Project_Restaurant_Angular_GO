import { TestBed } from '@angular/core/testing';
import { AppStore } from './app.store';

describe('AppStore', () => {
  let appStore: any;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AppStore],
    });

    appStore = TestBed.inject(AppStore);
  })

  it('should be created', () => {
    // Arrange Act Assert
    expect(appStore).toBeTruthy();
  });

  it('should have the initial showHeader value as false', () => {
    // Arrange Act Assert
    expect(appStore.showHeader()).toBeFalse();
  });

  it('should have the initial currentUrl value as an empty string', () => {
    // Arrange Act Assert
    expect(appStore.currentUrl()).toBe('');
  });

  describe('enableHeader()', () => {
    it('should set showHeader to true', () => {
      // Arrange Act
      appStore.enableHeader();

      // Assert
      expect(appStore.showHeader()).toBeTrue();
    });
  });

  describe('disableHeader()', () => {
    it('should set showHeader to false', () => {
      // Arrange
      appStore.enableHeader();

      //  Act
      appStore.disableHeader();

      // Assert
      expect(appStore.showHeader()).toBeFalse();
    });
  });

  describe('setCurrentUrl()', () => {
    it('should update currentUrl to the provided value', () => {
      // Arrange Act Assert
      const newUrl = '/new-url';

      // Act
      appStore.setCurrentUrl(newUrl);

      // Assert
      expect(appStore.currentUrl()).toBe(newUrl);
    });

    it('should update currentUrl to an empty string when called with an empty string', () => {
      // Arrange
      appStore.setCurrentUrl('newUrl');
      const newUrl = '';

      // Act
      appStore.setCurrentUrl(newUrl);

      // Assert
      expect(appStore.currentUrl()).toBe(newUrl);
    });
  });
});
