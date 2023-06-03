import { NgModule } from '@angular/core';
import { AppComponent } from './app.component';
import { appDeclarations } from './app-declarations';
import { appImports } from './app-imports';
import { appProviders } from './app-provider';

@NgModule({
  declarations: [
    ...appDeclarations,
  ],
  imports: [
    ...appImports
  ],
  providers: [
    ...appProviders
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
