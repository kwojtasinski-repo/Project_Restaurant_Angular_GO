import { changeInputValue, changeSelectIndex } from 'src/app/unit-test-fixtures/dom-utils';

export class ProductForm {
  private readonly nativeElement: any;

  public constructor(nativeElement: any) {
    this.nativeElement = nativeElement;
  }

  getProductForm(): any {
    return this.nativeElement?.querySelector('form');
  }

  getProductNameInput(): any {
    return this.nativeElement?.querySelector('#product-name');
  }

  getProductCostInput(): any {
    return this.nativeElement?.querySelector('#product-cost');
  }

  getProductDescriptionInput(): any {
    return this.nativeElement?.querySelector('#product-description');
  }

  getProductCategorySelectList(): any {
    return this.nativeElement?.querySelector('#product-category');
  }
  
  changeProductForm(productName: any, productDescription: any, productCost: any, selectedCategory: any) {
    changeInputValue(this.getProductNameInput(), productName);
    changeInputValue(this.getProductDescriptionInput(), productDescription);
    changeInputValue(this.getProductCostInput(), productCost);
    changeSelectIndex(this.getProductCategorySelectList(), selectedCategory);
  }

  getProductFormErrors(): any[] {
    return Array.from(this.nativeElement.querySelectorAll('.invalid-feedback')) as any[];
  }
}
