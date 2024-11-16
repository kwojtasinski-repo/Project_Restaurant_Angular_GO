import { Component, OnDestroy, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { getError } from 'src/app/stores/category/category.selectors';
import { CategoryState } from 'src/app/stores/category/category.state';
import { take, tap, finalize, catchError, EMPTY, Observable, shareReplay, BehaviorSubject } from 'rxjs';
import * as CategoryActions from 'src/app/stores/category/category.actions';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { ActivatedRoute } from '@angular/router';
import { NgxSpinnerService } from 'ngx-spinner';

@Component({
  selector: 'app-edit-category',
  templateUrl: './edit-category.component.html',
  styleUrls: ['./edit-category.component.scss']
})
export class EditCategoryComponent implements OnInit, OnDestroy {
  public category$: Observable<Category | undefined> | undefined;
  public isLoading$: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
  public error$ = this.store.select(getError);
  public error: string | undefined;

  constructor(private store: Store<CategoryState>, private categoryService: CategoryService, private route: ActivatedRoute, private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id') ?? '';
    this.category$ = this.categoryService.get(id)
      .pipe(
        take(1),
        shareReplay(),
        tap(() => {
          this.isLoading$.next(true);
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading$.next(false);
          this.spinnerService.hide();
        }),
        catchError((error) => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
          return EMPTY;
        })
      );
  }

  public onCategoryChange(category: Category): void {
    this.store.dispatch(CategoryActions.categoryFormUpdate({
      category
    }));
  }

  public onSubmit(): void {
    this.store.dispatch(CategoryActions.categoryUpdateRequestBegin());
  }

  public cancelClick(): void {
    this.store.dispatch(CategoryActions.categoryCancelOperation());
  }

  public ngOnDestroy(): void {
    this.store.dispatch(CategoryActions.clearErrors());
  }

  public show(obj: any): string {
    return JSON.stringify(obj);
  }
}
