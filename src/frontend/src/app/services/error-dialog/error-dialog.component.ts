import { Component } from '@angular/core';
import { NgIf } from '@angular/common';

@Component({
    selector: 'app-error-dialog',
    templateUrl: './error-dialog.component.html',
    styleUrls: ['./error-dialog.component.scss'],
    standalone: true,
    imports: [NgIf]
})
export class ErrorDialogComponent {
  public message = '';
  public status: number | undefined;
  
  public onClose(): void {

  }
}
