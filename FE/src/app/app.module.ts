import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { RepoComponent } from './repo/repo.component';
import { FilesComponent } from './files/files.component';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { RepoEffects } from './repo.effects';
import { RepoReducer } from './repo.reducer';

@NgModule({
  declarations: [AppComponent, RepoComponent, FilesComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserModule,
    FormsModule,
    HttpClientModule,
    StoreModule.forRoot({ repo: RepoReducer }),
    EffectsModule.forRoot([RepoEffects]),
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
