import { Component, EventEmitter, Input, Output } from '@angular/core';

@Component({
  selector: 'app-search-bar',
  templateUrl: './search-bar.component.html',
  styleUrls: ['./search-bar.component.scss']
})
export class SearchBarComponent {
  @Input()
  public term: string = '';

  @Input()
  public placeholder: string = '';

  @Output()
  public termChange: EventEmitter<string> = new EventEmitter<string>();

  @Output()
  public onChange: EventEmitter<string> = new EventEmitter<string>();

  public search(): void {
    this.onChange.emit(this.term);
  }
}
