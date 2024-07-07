import { Component, OnInit } from '@angular/core';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { finalize, take, tap } from 'rxjs';
import { NgxSpinnerService } from 'ngx-spinner';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent implements OnInit {
  public categories: Category[] = [];
  public categoriesToShow: Category[] = [];
  public term: string = '';
  public isLoading: boolean = true;
  public error: string | undefined;

  constructor(private categoryService: CategoryService, private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    this.categoryService.getAll()
      .pipe(
        take(1),
        tap(() => {
          this.isLoading = true;
          this.spinnerService.show();
        },
        finalize(() => {
          this.isLoading = false;
          this.spinnerService.hide();
        })
      )
      )
      .subscribe({ next: c => {
          this.categories = c;
          this.categoriesToShow = c;
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          console.error(error);
        }
      });
  }

  public search(term: string): void {
    this.categoriesToShow = this.categories.filter(c => c.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }
}
