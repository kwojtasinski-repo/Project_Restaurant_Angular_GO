import { inject, Injectable, Signal, signal, WritableSignal } from '@angular/core';
import { Router } from '@angular/router';
import { catchError, EMPTY, exhaustMap, Observable, take } from 'rxjs';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';

@Injectable({
    providedIn: 'root'
})
export class CategoryStore {
  private readonly _category: WritableSignal<Category | null> = signal(null);
  private readonly _error: WritableSignal<string | null> = signal(null);

  private router = inject(Router);
  private categoryService = inject(CategoryService);
  
  public get category(): Signal<Category | null> {
    return this._category.asReadonly();
  }

  public get error(): Signal<string | null> {
    return this._error.asReadonly();
  }

  public addCategory(): void {
    const category = this._category();
    if (!category) {
      this._error.set('Kategoria nie może być pusta');
      return;
    }

    this.clearErrors();
    this.categoryService.add(category)
      .pipe(
        take(1),
        catchError((error, caught) => this.handleBackendErrors(error, caught)),
        exhaustMap(() => this.addCategorySuccess())
      ).subscribe();
  }

  public updateCategory(): void {
    const category = this._category();
    if (!category || !category.id) {
      this._error.set('Kategoria nie może być pusta');
      return;
    }

    this.clearErrors();
    this.categoryService.update(category)
      .pipe(
        take(1),
        catchError((error, caught) => this.handleBackendErrors(error, caught)),
        exhaustMap(() => this.updateCategorySuccess())
      ).subscribe();
  }

  public updateCategoryForm(category: Category): void {
    this._category.set(category);
  }

  public clearErrors(): void {
    this._error.set(null);
  }

  public clearCategoryForm(): void {
    this._category.set(null);
    this.clearErrors();
  }
  
  public cancelCategoryOperation(): void {
    this.router.navigate(['/categories']);
    this.clearCategoryForm();
  }

  private addCategorySuccess(): Observable<void> {
    this.router.navigate(['/categories']);
    this.clearCategoryForm();
    return EMPTY;
  }

  private updateCategorySuccess(): Observable<void> {
    this.router.navigate(['/categories']);
    this.clearCategoryForm();
    return EMPTY;
  }

  private handleBackendErrors(err: any, _: Observable<void>): Observable<void> {
    if (err.status === 0) {
      this._error.set('Sprawdź połączenie z internetem');
      return EMPTY;
    } else if (err.status >= 500) {
      this._error.set('Coś poszło nie tak, spróbuj później');
      return EMPTY;
    }
    
    this._error.set(err?.error?.errors);
    return EMPTY;
  }
}
