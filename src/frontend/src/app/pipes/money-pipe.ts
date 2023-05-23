import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
    name: 'money'
})
export class MoneyPipe implements PipeTransform {
    public transform(val: any): string {
        return new Intl.NumberFormat('pl-PL', {
            minimumFractionDigits: 2,
            maximumFractionDigits: 2,
        }).format(Number(val));
    }
}
