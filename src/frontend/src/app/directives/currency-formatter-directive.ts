import { Directive, HostListener, OnDestroy, Input, OnInit, AfterViewInit, inject } from '@angular/core';
import { NgControl } from '@angular/forms';
import { Subject, debounceTime, takeUntil } from 'rxjs';

@Directive({
    selector: '[currencyFormatter]',
    standalone: true
})
  export class CurrencyFormatterDirective implements OnDestroy, OnInit, AfterViewInit {
    private ngControl = inject(NgControl, { self: true });

    @Input() locale: string = 'en-US';
    @Input() minimumFractionDigits: number = 0;
    @Input() maximumFractionDigits: number = 2;
    
    private formatter: Intl.NumberFormat;
    private destroy$ = new Subject();
    private comma: string = '.';
  
    constructor() {
      this.formatter = new Intl.NumberFormat(this.locale, { minimumFractionDigits: this.minimumFractionDigits, maximumFractionDigits: this.maximumFractionDigits });
    }

    public ngOnInit(): void {
        if (this.minimumFractionDigits > this.maximumFractionDigits || this.maximumFractionDigits < this.minimumFractionDigits) {
            throw new Error(`Invalid minimumFractionDigits '${this.minimumFractionDigits}' and maximumFractionDigits '${this.maximumFractionDigits}'`);
        }

        this.formatter = new Intl.NumberFormat(this.locale, { minimumFractionDigits: this.minimumFractionDigits, maximumFractionDigits: this.maximumFractionDigits });
        this.comma = this.formatter.format(0.1).charAt(1);
    }
  
    public ngAfterViewInit() {
      this.setValue(this.formatPrice(this.ngControl.value))
      this.ngControl.control?.valueChanges
        .pipe(debounceTime(10), takeUntil(this.destroy$))
        .subscribe(this.updateValue.bind(this));      
    }
  
    public ngOnDestroy() {
      this.setValue(this.unformatValue(this.ngControl.value));
      this.destroy$.complete();
    }
  
    @HostListener('focus') onFocus() {
      this.setValue(this.unformatValue(this.ngControl.value));
    }
  
    @HostListener('blur') onBlur() {
      const value = this.getNumberFromValue(this.ngControl.value || '0.00');
      !!value && this.setValue(this.formatPrice(value));
    }
  
    private updateValue(value: any) {
      const inputVal = value || '';
      const pattern = new RegExp('[^0-9' + this.comma + ']', 'g');
      this.setValue(inputVal ?
        this.validateDecimalValue(inputVal.replace(pattern, '')) : '');
    }
  
    private formatPrice(value: any) {
      return this.formatter.format(value);
    }
  
    private unformatValue(value: any) {
      const newValue = value.replace(' ', '');
      const separateThousand = ',';
      if (this.comma == separateThousand) {
        return newValue;
      }
      const pattern = new RegExp(separateThousand, 'g');
      return newValue.replace(pattern, '');
    }
  
    private validateDecimalValue(value: any) {
      const newValue = this.getNumberFromValue(value);

      if (Number.isNaN(Number(newValue))) {
        const strippedValue = newValue.slice(0, newValue.length - 1);
        return Number.isNaN(Number(strippedValue)) ? '' : strippedValue;
      }
      return value;
    }
  
    private setValue(value: any) {
      this.ngControl.control?.setValue(value, { emitEvent: false })
    }

    private getNumberFromValue(value: any) {
      return this.comma === ',' ? value.replace(',', '.') : value
    }
}
