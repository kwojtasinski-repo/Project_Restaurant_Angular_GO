import { Router } from '@angular/router';
import { CategoryStore } from './category.store';
import { TestBed } from '@angular/core/testing';
import { CategoryService } from 'src/app/services/category.service';
import { Category } from 'src/app/models/category';
import { throwError } from 'rxjs';
import { TestSharedModule } from 'src/app/unit-test-fixtures/test-share-module';
import { completeObservable } from 'src/app/unit-test-fixtures/observable-utils';

describe('CategoryStore', () => {
    let categoryStore: CategoryStore;
    let mockRouter: jasmine.SpyObj<Router>;
    let categoryService: CategoryService;
  
    beforeEach(() => {
      mockRouter = jasmine.createSpyObj('Router', ['navigate']);
  
      TestBed.configureTestingModule({
        imports: [TestSharedModule],
        providers: [
          CategoryStore,
          { provide: Router, useValue: mockRouter }
        ],
      });
  
      categoryService = TestBed.inject(CategoryService);
      categoryStore = TestBed.inject(CategoryStore);
    });

    describe('addCategory', () => {
      it('should set error if category is null', () => {
        categoryStore.addCategory();
        expect(categoryStore.error()!).toBe('Kategoria nie może być pusta');
      });
    
      it('should clear errors and call categoryService.add if category is valid', () => {
        const category: Category = { id: '1', name: 'Test Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'add').and.returnValue(completeObservable())

        categoryStore.addCategory();

        expect(categoryService.add).toHaveBeenCalledWith(category);
        expect(categoryStore.error()).toBeNull();
        expect(mockRouter.navigate).toHaveBeenCalledWith(['/categories']);
      });
    
      it('should handle backend error 500 correctly', () => {
        const category: Category = { id: '1', name: 'Test Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'add').and.returnValue(throwError(() => ({ status: 500 })));

        categoryStore.addCategory();

        expect(categoryStore.error()).toBe('Coś poszło nie tak, spróbuj później');
      });
    
      it('should handle internet problem connection correctly', () => {
        const category: Category = { id: '1', name: 'Test Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'add').and.returnValue(throwError(() => ({ status: 0 })));

        categoryStore.addCategory();

        expect(categoryStore.error()).toBe('Sprawdź połączenie z internetem');
      });
    });

    describe('updateCategory', () => {
      it('should set error if category is null or has no ID', () => {
        categoryStore.updateCategory();
        expect(categoryStore.error()!).toBe('Kategoria nie może być pusta');
    
        categoryStore.updateCategoryForm({ name: 'No ID' } as Category);
        categoryStore.updateCategory();
        expect(categoryStore.error()!).toBe('Kategoria nie może być pusta');
      });
    
      it('should call categoryService.update if category is valid', () => {
        const category: Category = { id: '2', name: 'Updated Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'update').and.returnValue(completeObservable());
    
        categoryStore.updateCategory();
    
        expect(categoryService.update).toHaveBeenCalledWith(category);
        expect(mockRouter.navigate).toHaveBeenCalledWith(['/categories']);
      });
    
      it('should handle backend error 500 correctly', () => {
        const category: Category = { id: '2', name: 'Updated Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'update').and.returnValue(throwError(() => ({ status: 500 })));
    
        categoryStore.updateCategory();
    
        expect(categoryStore.error()).toBe('Coś poszło nie tak, spróbuj później');
      });
    
      it('should handle internet problem connection correctly', () => {
        const category: Category = { id: '2', name: 'Updated Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'update').and.returnValue(throwError(() => ({ status: 0 })));
    
        categoryStore.updateCategory();
    
        expect(categoryStore.error()).toBe('Sprawdź połączenie z internetem');
      });
    });
    
    describe('updateCategoryForm', () => {
      it('should set the category correctly', () => {
        const category: Category = { id: '3', name: 'New Category', deleted: false };
        categoryStore.updateCategoryForm(category);
    
        expect(categoryStore.category()).toEqual(category);
      });
    });
    
    describe('clearErrors', () => {
      it('should clear errors', () => {
        spyOn(categoryService, 'update').and.returnValue(throwError(() => ({ status: 0 })));
        categoryStore.updateCategory();
        const error = categoryStore.error();

        categoryStore.clearErrors();
    
        const errorAfterClear = categoryStore.error();
        expect(error).toBeTruthy();
        expect(errorAfterClear).not.toBe(error);
        expect(errorAfterClear).toBeNull();
      });
    });
    
    describe('clearCategoryForm', () => {
      it('should clear the category and errors', () => {
        const category: Category = { id: '4', name: 'Temporary Category', deleted: false };
        categoryStore.updateCategoryForm(category);
        spyOn(categoryService, 'update').and.returnValue(throwError(() => ({ status: 0 })));
        categoryStore.updateCategory();
        const error = categoryStore.error();
    
        categoryStore.clearCategoryForm();
    
        expect(categoryStore.category()).not.toEqual(category);
        expect(categoryStore.error()).not.toEqual(error);
        expect(categoryStore.category()).toBeNull();
        expect(categoryStore.error()).toBeNull();
      });
    });
    
    describe('cancelCategoryOperation', () => {
      it('should navigate to /categories and clear the form', () => {
        const category: Category = { id: '5', name: 'Cancelable Category', deleted: false };
        categoryStore.updateCategoryForm(category);
    
        categoryStore.cancelCategoryOperation();
    
        expect(mockRouter.navigate).toHaveBeenCalledWith(['/categories']);
        expect(categoryStore.category()).toBeNull();
      });
    });
});
