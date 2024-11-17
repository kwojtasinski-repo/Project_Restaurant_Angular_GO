import { Component } from '@angular/core';


@Component({
    selector: 'app-error-dialog',
    templateUrl: './error-dialog.component.html',
    styleUrls: ['./error-dialog.component.scss'],
    standalone: true,
    imports: []
})
export class ErrorDialogComponent {
  public message = '';
  public status: number | undefined;
  
  public onClose(): void {

  }
}
