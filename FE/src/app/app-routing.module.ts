import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RepoComponent } from './repo/repo.component';
import { FilesComponent } from './files/files.component';

const routes: Routes = [
  { path: 'repo', component: RepoComponent },
  { path: 'files', component: FilesComponent },
  { path: '', redirectTo: '/repo', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
