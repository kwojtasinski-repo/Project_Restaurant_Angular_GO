import { ComponentFixture, TestBed } from '@angular/core/testing';
import { initialState } from 'src/app/stores/register/register.reducers';
import { provideMockStore } from '@ngrx/store/testing';

import { RegisterComponent } from './register.component';
import { RouterModule } from '@angular/router';
import { provideMockActions } from '@ngrx/effects/testing';
import { Actions } from '@ngrx/effects';
import { SpinnerButtonComponent } from '../spinner-button/spinner-button.component';
import { ReactiveFormsModule } from '@angular/forms';

describe('RegisterComponent', () => {
  let component: RegisterComponent;
  let fixture: ComponentFixture<RegisterComponent>;
  let actions: Actions;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        RegisterComponent,
        SpinnerButtonComponent
      ],
      imports: [
        ReactiveFormsModule,
        RouterModule.forRoot([]),
      ],
      providers: [
        provideMockStore({ initialState }),
        provideMockActions(() => actions)
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
