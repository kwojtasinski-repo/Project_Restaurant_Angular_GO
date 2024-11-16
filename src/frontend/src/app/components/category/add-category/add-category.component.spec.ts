import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddCategoryComponent } from './add-category.component';
import { provideMockStore } from '@ngrx/store/testing';
import { initialState } from 'src/app/stores/category/category.reducers';
import { ReactiveFormsModule } from '@angular/forms';
import { CategoryFormComponent } from '../category-form/category-form.component';

describe('AddCategoryComponent', () => {
  let component: AddCategoryComponent;
  let fixture: ComponentFixture<AddCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddCategoryComponent, CategoryFormComponent ],
      imports: [
        ReactiveFormsModule
      ],
      providers: [
        provideMockStore({ initialState })
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddCategoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
