import { Injectable } from '@angular/core';
import { BsModalRef, BsModalService } from 'ngx-bootstrap/modal';
import { ErrorDialogComponent } from './error-dialog/error-dialog/error-dialog.component';

@Injectable({
  providedIn: 'root'
})
export class ErrorDialogService {
  private opened = false;
  private modalRef: BsModalRef = new BsModalRef();

  constructor(private modalService: BsModalService) { }

  public openDialog(message: string, status?: number): void {
    if (!this.opened) {
      this.opened = true;
      this.modalRef = this.modalService.show(ErrorDialogComponent, {
        ignoreBackdropClick: true
      });
      this.modalRef.content.message = message;
      this.modalRef.content.status = status;
      this.modalRef.content.onClose = this.modalRef.hide;

      this.modalRef.onHidden?.subscribe(() => {
        this.opened = false;
      });
    }
  }
}
