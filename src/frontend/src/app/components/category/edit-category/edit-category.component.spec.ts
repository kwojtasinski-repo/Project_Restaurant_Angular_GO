import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditCategoryComponent } from './edit-category.component';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from 'src/app/stores/category/category.reducers';
import { NgxSpinnerModule } from 'ngx-spinner';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { provideRouter } from '@angular/router';
import { CategoryFormComponent } from '../category-form/category-form.component';

describe('EditCategoryComponent', () => {
  let component: EditCategoryComponent;
  let fixture: ComponentFixture<EditCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
    providers: [
        provideRouter([]),
        provideMockStore({ initialState }),
        {
            provide: 'API_URL', useValue: ''
        }
    ],
    imports: [
        NgxSpinnerModule,
        ReactiveFormsModule,
        HttpClientModule,
        EditCategoryComponent, CategoryFormComponent
    ]
})
    .compileComponents();

    fixture = TestBed.createComponent(EditCategoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
