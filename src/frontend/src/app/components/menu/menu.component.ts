import { Component } from '@angular/core';
import { Product } from 'src/app/models/product';
import { ProductService } from 'src/app/services/product.service';
import { Observable } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-menu',
  templateUrl: './menu.component.html',
  styleUrls: ['./menu.component.scss']
})
export class MenuComponent {
  public user$ = this.authService.getUser();

  constructor(private productService: ProductService, private authService: AuthService) { }

  public showProducts(): Observable<Product[]> {
    return this.productService.getAll();
  }
}
