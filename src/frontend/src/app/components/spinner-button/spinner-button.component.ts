import { Component, EventEmitter, Input, Output } from '@angular/core';
import { SpinnerVersion } from './spinner-version';


@Component({
    selector: 'app-spinner-button',
    templateUrl: './spinner-button.component.html',
    styleUrls: ['./spinner-button.component.scss'],
    standalone: true,
    imports: []
})
export class SpinnerButtonComponent {
  @Input()
  public buttonText: string = '';

  @Input()
  public className: string = '';

  @Input()
  public disabled: boolean = false;

  @Input()
  public version: SpinnerVersion = SpinnerVersion.classic;
  
  @Output() 
  buttonClick: EventEmitter<any> = new EventEmitter();

  public click(event: any): void {
    this.buttonClick.emit(event);
  }
}
