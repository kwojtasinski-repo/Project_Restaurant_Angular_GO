import { resolveApiUrl } from "./providers/api-url-provider";

export const appProviders = [
    {
      provide: "API_URL", 
      useFactory: resolveApiUrl,
      multi: true
    },
]
