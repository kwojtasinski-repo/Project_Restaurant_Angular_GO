import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
    selector: 'app-search-bar',
    templateUrl: './search-bar.component.html',
    styleUrls: ['./search-bar.component.scss'],
    standalone: true,
    imports: [FormsModule]
})
export class SearchBarComponent {
  @Input()
  public term: string = '';

  @Input()
  public placeholder: string = '';

  @Output()
  public termChange: EventEmitter<string> = new EventEmitter<string>();

  public search(): void {
    this.termChange.emit(this.term);
  }
}
