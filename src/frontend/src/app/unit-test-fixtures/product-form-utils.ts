export function getProductForm(nativeElement: any): any {
    return nativeElement?.querySelector('form');
}

export function getProductNameInput(nativeElement: any): any {
    return nativeElement?.querySelector('#product-name');
}

export function getProductCostInput(nativeElement: any): any {
    return nativeElement?.querySelector('#product-cost');
}

export function getProductDescriptionInput(nativeElement: any): any {
    return nativeElement?.querySelector('#product-description');
}

export function getProductCategorySelectList(nativeElement: any): any {
    return nativeElement?.querySelector('#product-category');
}
