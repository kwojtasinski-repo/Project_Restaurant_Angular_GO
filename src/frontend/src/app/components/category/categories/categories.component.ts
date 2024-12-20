import { Component, OnInit, Signal, WritableSignal, computed, inject, signal } from '@angular/core';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { EMPTY, catchError, finalize, tap } from 'rxjs';
import { NgxSpinnerService } from 'ngx-spinner';
import { SearchBarComponent } from '../../search-bar/search-bar.component';
import { RouterLink } from '@angular/router';

@Component({
    selector: 'app-categories',
    templateUrl: './categories.component.html',
    styleUrls: ['./categories.component.scss'],
    standalone: true,
    imports: [RouterLink, SearchBarComponent]
})
export class CategoriesComponent implements OnInit {
  private readonly categoryService = inject(CategoryService);
  private readonly spinnerService = inject(NgxSpinnerService);

  public readonly categories: WritableSignal<Category[]> = signal([]);
  public readonly term: WritableSignal<string> = signal('');
  public readonly isLoading: WritableSignal<boolean> = signal(true);
  public readonly error: WritableSignal<string | undefined> = signal<string | undefined>(undefined);

  // Derived signal for filtered categories
  public readonly categoriesToShow: Signal<Category[]> = computed(() => {
    const term = this.term();
    return this.categories().filter(c =>
      c.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase())
    );
  });

  public ngOnInit(): void {
    this.categoryService.getAll()
      .pipe(
        tap(() => {
          this.isLoading.set(true);
          this.spinnerService.show();
        }),
        finalize(() => {
          this.isLoading.set(false);
          this.spinnerService.hide();
        }),
        catchError((error) => {
          if (error.status === 0) {
            this.error.set('Sprawdź połączenie z internetem');
          } else if (error.status === 500) {
            this.error.set('Coś poszło nie tak, spróbuj ponownie później');
          }
          console.error(error);
          return EMPTY;
        })
      ).subscribe(categories => this.categories.set(categories));
  }
}
