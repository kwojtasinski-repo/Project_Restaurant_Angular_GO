import { Directive, OnDestroy, OnInit, inject } from '@angular/core';
import { FormGroupDirective } from '@angular/forms';
import { LoginState } from '../stores/login/login.state';
import { Store } from '@ngrx/store';
import { Subscription, debounceTime, take } from 'rxjs';
import { loginFormUpdate } from '../stores/login/login.actions';

@Directive({
    selector: '[loginForm]',
    standalone: true
})
export class LoginFormDirective implements OnInit, OnDestroy {
    private formGroupDirective = inject(FormGroupDirective);
    private store = inject<Store<LoginState>>(Store);

    public formChange : Subscription = new Subscription();
    
    public ngOnInit() {
        // Update the form value based on the state
        this.store.select(state => state.credentials)
            .pipe(take(1))
            .subscribe(formValue => {
                this.formGroupDirective.form.patchValue(formValue);
            });
        
        this.formChange = this.formGroupDirective.form.valueChanges
            .pipe(debounceTime(10))
            .subscribe(value => {
                this.store.dispatch(loginFormUpdate({ credentials: {
                    email: value.emailAddress,
                    password: value.password
                }}));
            });
    }

    public ngOnDestroy(): void {
        this.formChange.unsubscribe();
    }
}
