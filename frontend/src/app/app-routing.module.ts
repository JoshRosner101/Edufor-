import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router'; // CLI imports router
import { ByIDComponent } from './by-id/by-id.component';
import { MainPageComponent } from './main-page/main-page.component';

const routes: Routes = [
  { path: 'home/:id', component: ByIDComponent },
  { path: 'home', component: MainPageComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full'},
];

// configures NgModule imports and exports
@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }