import { Injectable } from '@angular/core';
import { MatLegacySnackBar as MatSnackBar } from '@angular/material/legacy-snack-bar';
import { SwUpdate } from '@angular/service-worker';

@Injectable({
  providedIn: 'root',
})
export class UpdateService {
  constructor(private swUpdate: SwUpdate, snackbar: MatSnackBar) {
    this.swUpdate.available.subscribe((evt) => {
      const snack = snackbar.open('Update Available', 'Reload');

      snack.onAction().subscribe(() => {
        window.location.reload();
      });
    });
  }
}
