import { TestBed } from '@angular/core/testing';
import { take } from 'rxjs';

import { AuthenticationService } from './authentication.service';
import { HttpClient } from '@angular/common/http';
import { completeObservable, errorObservable } from '../unit-test-fixtures/observable-utils';
import { User } from '../models/user';
import { TestSharedModule } from '../unit-test-fixtures/test-share-module';
import { createUser } from '../unit-test-fixtures/user-utils';

describe('AuthenticationServiceService', () => {
  let service: AuthenticationService;
  let httpClient: HttpClient

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [TestSharedModule]
    }).compileComponents();

    service = TestBed.inject(AuthenticationService);
    httpClient = TestBed.inject(HttpClient);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should login', () => {
    const user = createUser();
    spyOn(httpClient, 'post').and.returnValue(completeObservable<void>());
    spyOn(httpClient, 'get').and.returnValue(completeObservable<User>(user));

    service.login({ email: 'email@email.com', password: '' })
      .pipe(take(1))
      .subscribe(u => expect(u).toEqual(user));
  });

  it('login while some error occured after getting user info should return error action', () => {
    const error = 'error occured'
    spyOn(httpClient, 'post').and.returnValue(completeObservable<void>());
    spyOn(httpClient, 'get').and.returnValue(errorObservable(error));

    service.login({ email: 'email@email.com', password: '' })
      .pipe(take(1))
      .subscribe({
        next: val => expect(val).toBeNull(),
        error: e => {
          expect(e).not.toBeUndefined();
          expect(e).not.toBeNull();
          expect(e).toEqual(error);
        }
      });
  });

  it('login while some error occured after getting user info should return error action', () => {
    const error = 'error occured'
    spyOn(httpClient, 'post').and.returnValue(errorObservable(error));

    service.login({ email: 'email@email.com', password: '' })
      .pipe(take(1))
      .subscribe({
        next: val => expect(val).toBeNull(),
        error: e => {
          expect(e).not.toBeUndefined();
          expect(e).not.toBeNull();
          expect(e).toEqual(error);
        }
      });
  });
});
