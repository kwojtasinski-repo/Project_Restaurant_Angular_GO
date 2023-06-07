import { Component } from '@angular/core';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { take } from 'rxjs';
import { NgxSpinnerService } from 'ngx-spinner';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent {
  public categories: Category[] = [];
  public categoriesToShow: Category[] = [];
  public term: string = '';
  public error: string | undefined;

  constructor(private categoryService: CategoryService, private spinnerService: NgxSpinnerService) { }

  public ngOnInit(): void {
    this.spinnerService.show();
    this.categoryService.getAll()
      .pipe(take(1))
      .subscribe({ next: c => {
          this.categories = c;
          this.categoriesToShow = c;
          this.spinnerService.hide();
        }, error: error => {
          if (error.status === 0) {
            this.error = 'Sprawdź połączenie z internetem';
          } else if (error.status === 500) {
            this.error = 'Coś poszło nie tak, spróbuj ponownie później';
          }
          this.spinnerService.hide();
          console.error(error);
        }
      });
  }

  public search(term: string): void {
    this.categoriesToShow = this.categories.filter(c => c.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }
}
