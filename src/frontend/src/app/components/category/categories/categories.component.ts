import { Component } from '@angular/core';
import { Category } from 'src/app/models/category';
import { CategoryService } from 'src/app/services/category.service';
import { take } from 'rxjs';

@Component({
  selector: 'app-categories',
  templateUrl: './categories.component.html',
  styleUrls: ['./categories.component.scss']
})
export class CategoriesComponent {
  public categories: Category[] = [];
  public categoriesToShow: Category[] = [];
  public term: string = '';

  constructor(private categoryService: CategoryService) { }

  public ngOnInit(): void {
    this.categoryService.getAll()
      .pipe(take(1))
      .subscribe(c => {
        this.categories = c;
        this.categoriesToShow = c;
      });
  }

  public search(term: string): void {
    this.categoriesToShow = this.categories.filter(c => c.name.toLocaleLowerCase().startsWith(term.toLocaleLowerCase()));
  }
}
