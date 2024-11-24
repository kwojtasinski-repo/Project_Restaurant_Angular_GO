import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditCategoryComponent } from './edit-category.component';
import { CategoryFormComponent } from '../category-form/category-form.component';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';

describe('EditCategoryComponent', () => {
  let component: EditCategoryComponent;
  let fixture: ComponentFixture<EditCategoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        EditCategoryComponent,
        CategoryFormComponent,
        TestSharedModule
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
