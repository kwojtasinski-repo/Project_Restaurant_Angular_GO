import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { customRoutes } from './routes';

const routes: Routes = customRoutes;

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
