import { Observable } from 'rxjs';

export function completeObservable<T>(value: T | undefined = undefined): Observable<T> {
    return new Observable<T>((o) => {
        if (value) {               
            o.next(value);
        } else {
            o.next();
        }
        o.complete();
    });
}

export function errorObservable<T>(error: any): Observable<T> {
    return new Observable<T>(o => o.error(error));
}
