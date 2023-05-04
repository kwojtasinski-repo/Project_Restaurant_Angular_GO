import { Directive, HostListener, OnDestroy, Self, Input, OnInit } from "@angular/core";
import { NgControl } from "@angular/forms";
import { Subject, takeUntil } from "rxjs";

@Directive({
    selector: '[currencyFormatter]'
  })
  export class CurrencyFormatterDirective implements OnDestroy, OnInit {
    @Input() locale: string = 'en-US';
    @Input() minimumFractionDigits: number = 0;
    @Input() maximumFractionDigits: number = 2;
    
    private formatter: Intl.NumberFormat;
    private destroy$ = new Subject();
    private comma: string = '.';
  
    constructor(@Self() private ngControl: NgControl) {
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
        .pipe(takeUntil(this.destroy$))
        .subscribe(this.updateValue.bind(this));      
    }
  
    private updateValue(value: any) {
      let inputVal = value || '';
      debugger
      const pattern = new RegExp("[^0-9" + this.comma + "]", 'g');
      this.setValue(!!inputVal ?
        this.validateDecimalValue(inputVal.replace(pattern, '')) : '');
    }
  
    @HostListener('focus') onFocus() {
      this.setValue(this.unformatValue(this.ngControl.value));
    }
  
    @HostListener('blur') onBlur() {
      let value = this.ngControl.value || '';
      !!value && this.setValue(this.formatPrice(value));
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
      const newValue = this.comma === ',' ? value.replace(',', '.') : value;

      // Check to see if the value is a valid number or not
      if (Number.isNaN(Number(newValue))) {
        // strip out last char as this would have made the value invalid
        const strippedValue = newValue.slice(0, newValue.length - 1);
  
        // if value is still invalid, then this would be copy/paste scenario
        // and in such case we simply set the value to empty
        return Number.isNaN(Number(strippedValue)) ? '' : strippedValue;
      }
      return newValue;
    }
  
    private setValue(value: any) {
      this.ngControl.control?.setValue(value, { emitEvent: false })
    }
  
    public ngOnDestroy() {
      this.setValue(this.unformatValue(this.ngControl.value));
      this.destroy$.complete();
    }
}
