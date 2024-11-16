import { ErrorHandler } from '@angular/core';
import { GlobalErrorHandler } from './global-error-handler';
import { resolveApiUrl } from './providers/api-url-provider';
import { BsModalService } from 'ngx-bootstrap/modal';

export const appProviders = [
    {
      provide: 'API_URL', 
      useFactory: resolveApiUrl,
      multi: true
    },
    BsModalService,
    {
      // processes all errors
      provide: ErrorHandler,
      useClass: GlobalErrorHandler,
    },
]
