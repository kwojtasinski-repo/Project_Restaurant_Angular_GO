import { ActivatedRoute, convertToParamMap } from '@angular/router';

export function createActivatedRouteProvider(routeObj: any) {
    return {
        provide: ActivatedRoute,
          useValue: {
              snapshot: {
                  paramMap: convertToParamMap({
                    ...routeObj
                  }),
              },
          },
    }
}
