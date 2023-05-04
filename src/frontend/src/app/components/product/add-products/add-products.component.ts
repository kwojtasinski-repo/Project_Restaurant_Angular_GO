import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-add-products',
  templateUrl: './add-products.component.html',
  styleUrls: ['./add-products.component.scss']
})
export class AddProductsComponent {
  public addProduct: FormGroup;
  
  constructor() {
    this.addProduct = new FormGroup({
      productName: new FormControl('', Validators.compose([Validators.required, Validators.maxLength(100), Validators.minLength(3)])),
      productDescription: new FormControl('', Validators.maxLength(5000)),
      productCost: new FormControl('0.00', Validators.compose([Validators.required, Validators.min(0)])),
      productCategory: new FormControl('', Validators.required),
    });
  }

  public onSubmit() {
    console.log(this.addProduct);
  }

  public onPriceChange(value: string) {
    console.log(new Number(value.replace(/[^0-9.-]+/g,"")));
  }
}
