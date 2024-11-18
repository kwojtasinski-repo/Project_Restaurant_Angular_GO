import { ErrorHandler, Injectable, Injector, NgZone } from '@angular/core';
import { ErrorDialogService } from './services/error-dialog.service';
import { NgxSpinnerService } from 'ngx-spinner';

@Injectable()
export class GlobalErrorHandler implements ErrorHandler {
  constructor(
    private injector: Injector,
  ) {}

  handleError(error: any) {
    const ngZone = this.injector.get(NgZone);
    const errorDialogService = this.injector.get(ErrorDialogService);
    this.injector.get(NgxSpinnerService).hide();

    ngZone.run(() =>
      errorDialogService.openDialog(
        error?.message || 'Undefined client error',
        error?.status
      )
    );

    console.error('Error from global error handler', error);
  }
}